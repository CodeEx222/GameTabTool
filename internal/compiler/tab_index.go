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
	// ##var	tableType	tableFileName	group	comment
	tempMode := make(map[int]string)
	for loadSheet.IsRowEmpty(startIndex, -1) == false {
		rowString := helper.ReadSheetRow(loadSheet, startIndex)

		if util.CheckStringType(rowString[0], "var") {
			// 这一行是类型定义
			for i, strVal := range rowString {
				if i == 0 {
					continue
				}

				if strVal == "tableType" {
					tempMode[i] = strVal
				} else if strVal == "tableFileName" {
					// 这个是表名
					tempMode[i] = strVal
				} else if strVal == "group" {
					// 这个是分组
					tempMode[i] = strVal
				} else if strVal == "comment" {
					// 这个是注释
					tempMode[i] = strVal
				}

			}

		} else if util.CheckValueType(rowString[0]) {
			// 值
			var tempCell model.IndexDefine
			for index, strValue := range rowString {
				checkType, isfind := tempMode[index]
				if !isfind {
					continue
				}
				if checkType == "tableType" {
					// 这个是类型
					tempCell.TableType = strValue
				} else if checkType == "tableFileName" {
					// 这个是表名
					tempCell.TableFilePath = strValue
					tempCell.TableFileName = strValue
				} else if checkType == "group" {
					// 这个是分组
					tempCell.Group = strings.Split(strValue, "|")
				} else if checkType == "comment" {
					// 这个是注释
					tempCell.Comment = strValue
				}
			}

			pragmaList = append(pragmaList, &tempCell)
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
	globals.IndexDefine = parseIndexRow(indexTabs, fileName)

	return nil
}
