package helper

import (
	"gametabtool/internal/report"
	"gametabtool/internal/util"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

type XlsxFile struct {
	// 读取excel对象
	file *excelize.File
	// excel文件中的所有sheet
	sheets []TableSheet
}

func (selfObj *XlsxFile) Sheets() (ret []TableSheet) {
	return selfObj.sheets
}

func (selfObj *XlsxFile) Save(filename string) error {
	return selfObj.file.SaveAs(filename)
}

func (selfObj *XlsxFile) Load(filename string) (err error) {

	var file *excelize.File

	file, err = excelize.OpenFile(filename)
	if err != nil {
		return err
	}

	selfObj.FromXFile(file)

	return nil
}

func (selfObj *XlsxFile) FromXFile(file *excelize.File) {
	selfObj.file = file

	SheetList := file.GetSheetList()
	for _, sheet := range SheetList {
		appobj, result := newXlsxSheet(sheet, file)
		if result {
			selfObj.sheets = append(selfObj.sheets, appobj)
		} else {
			report.ToolsLog.Info("  [%s] sheet is null, skip", sheet)
		}

	}
}

func NewXlsxFile() TableFile {

	self := &XlsxFile{}

	return self
}

type XlsxSheet struct {
	SheetName string
	ExcelFile *excelize.File
	MaxCol    int
	MaxRows   int
}

func (selfObj *XlsxSheet) Name() string {
	return selfObj.SheetName
}

func (selfObj *XlsxSheet) MaxColumn() int {
	return selfObj.MaxCol
}

func (selfObj *XlsxSheet) IsRowEmpty(row, maxCol int) bool {

	if maxCol == -1 {
		maxCol = selfObj.MaxCol
	}

	for col := 0; col < maxCol; col++ {

		data := selfObj.GetValue(row, col, nil)

		if data != "" {
			return false
		}
	}

	return true
}

func (selfObj *XlsxSheet) GetValue(row, col int, opt *ValueOption) (ret string) {
	ColStr := util.ConvertNumToChar(col + 1)
	ColStr = ColStr + strconv.Itoa(row+1)
	ValueStr, err := selfObj.ExcelFile.GetCellValue(selfObj.SheetName, ColStr)
	if err != nil {
		panic(err)
	}

	// 浮点数单元格按原样输出
	if opt != nil && opt.ValueAsFloat {
		ret, _ = util.GeneralNumericScientific(ValueStr, true)
		ret = strings.TrimSpace(ret)
	} else {
		// 取列头所在列和当前行交叉的单元格
		ret = strings.TrimSpace(ValueStr)
	}

	return
}

func newXlsxSheet(SheetName string, ExcelFile *excelize.File) (TableSheet, bool) {
	rows, err := ExcelFile.GetRows(SheetName)
	if err != nil {
		panic(err)
	}

	if len(rows) <= 0 {
		return nil, false
	}

	// err 可以忽略 前面 ExcelFile.GetRows 已经判断过了
	mergeCells, _ := ExcelFile.GetMergeCells(SheetName)

	MaxColValue := len(rows[0])

	for _, cellValue := range mergeCells {
		checkStr := cellValue[0]
		checkStrArray := strings.Split(checkStr, ":")

		checkStartResult, _, y1 := util.ConvertExcelCellToNumPos(checkStrArray[0])
		checkEndResult, x2, _ := util.ConvertExcelCellToNumPos(checkStrArray[1])

		if checkStartResult && checkEndResult {
			// 有合并单元格的情况下，取最大列数, 默认第一行的列数
			if y1 == 1 && x2 > MaxColValue {
				MaxColValue = x2
			}
		}

	}

	return &XlsxSheet{
		SheetName: SheetName,
		ExcelFile: ExcelFile,
		MaxCol:    MaxColValue,
		MaxRows:   len(rows),
	}, true
}
