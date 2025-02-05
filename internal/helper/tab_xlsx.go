package helper

import (
	"fmt"
	"gametabtool/internal/report"
	"gametabtool/internal/util"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

type XlsxFile struct {
	file *excelize.File

	sheets   []TableSheet
	cacheDir string
}

func (selfObj *XlsxFile) Sheets() (ret []TableSheet) {

	return selfObj.sheets
}

func (selfObj *XlsxFile) Save(filename string) error {
	return selfObj.file.SaveAs(filename)
}

func (selfObj *XlsxFile) Load(filename string) (err error) {

	var file *excelize.File

	if selfObj.cacheDir == "" {
		file, err = excelize.OpenFile(filename)
		if err != nil {
			return err
		}
	} else {
		panic("不支持cache")
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

func NewXlsxFile(cacheDir string) TableFile {

	self := &XlsxFile{
		cacheDir: cacheDir,
	}

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

	mergeCells, err := ExcelFile.GetMergeCells(SheetName)
	if err != nil {
		panic(err)
	}

	MaxColValue := len(rows[0])

	for _, cellValue := range mergeCells {
		checkStr := cellValue[0]
		checkStrArray := strings.Split(checkStr, ":")
		checkStrEnd := checkStrArray[1]

		checkStrAlp := ""
		checkStrNum := ""
		for _, c := range checkStrEnd {
			if c >= 'A' && c <= 'Z' {
				checkStrAlp += string(c)
			} else {
				checkStrNum += string(c)
			}
		}

		if checkStrNum == "1" {
			ColNum := util.ConvertNumToChar(checkStrAlp)
			if ColNum > MaxColValue {
				MaxColValue = ColNum
			}

		}

	}

	return &XlsxSheet{
		SheetName: SheetName,
		ExcelFile: ExcelFile,
		MaxCol:    len(rows[0]),
		MaxRows:   len(rows),
	}, true
}
