package compiler

import (
	"gametabtool/internal/helper"
	"gametabtool/internal/model"
	"gametabtool/internal/report"
	"gametabtool/internal/util"
	"path"
	"strings"
)

func parseIndexRow(tab helper.TableFile, fileName string) (pragmaList []*model.IndexDefine) {

	// Sheet1
	var loadSheet helper.TableSheet
	for _, sheet := range tab.Sheets() {
		if sheet.Name() == "Sheet1" {
			loadSheet = sheet
		}
	}

	if loadSheet == nil {
		report.LogTableError(report.UnknownSheet, fileName)
		return
	}

	startIndex := 0
	for loadSheet.IsRowEmpty(startIndex, -1) == false {
		rowString := helper.ReadSheetRow(loadSheet, startIndex)

		if util.CheckStringType(rowString[0], "var") {
			// 这一行是类型定义
			for i, s := range rowString {
				if i == 0 {
					continue
				}

			}

		} else {
			// 值
		}

		startIndex++
	}

	return
}

func LoadIndexTable(globals *model.Globals) error {

	if globals.IndexDataPath == "" {
		return nil
	}

	fileName := ""
	// 加载原始数据
	if len(globals.ExcelDataPath) > 0 {
		fileName = path.Join(globals.ExcelDataPath, globals.IndexDataPath)
	}

	loader := helper.NewFileLoader(true)
	indexTabs, err := loader.GetFile(fileName)

	if err != nil {
		return err
	}
	parseIndexRow(indexTabs, fileName)

	return nil
}
