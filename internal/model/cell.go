package model

import (
	"fmt"
	"gametabtool/internal/util"
)

type Cell struct {
	Value     string
	ValueList []string // merge之后, 数组值保存在这里
	Row       int      // base 0
	Col       int      // base 0
	Table     *DataTable
}

// 全拷贝
func (selfObj *Cell) CopyFrom(c *Cell) {
	selfObj.Value = c.Value
	selfObj.Row = c.Row
	selfObj.Col = c.Col
	selfObj.Table = c.Table
}

func (selfObj *Cell) String() string {

	var file, sheet string
	if selfObj.Table != nil {
		file = selfObj.Table.FileName
		sheet = selfObj.Table.SheetName
	}

	var value string
	if len(selfObj.ValueList) > 0 {
		value = fmt.Sprintf("%+v", selfObj.ValueList)
	} else {
		value = selfObj.Value
	}

	return fmt.Sprintf("'%s' @%s|%s(%s)", value, file, sheet, util.R1C1ToA1(selfObj.Row+1, selfObj.Col+1))
}
