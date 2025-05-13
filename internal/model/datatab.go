package model

import (
	"fmt"
	"strings"
)

// DataTable 表格的完整数据，表头有屏蔽时，对应行值为空
type DataTable struct {
	HeaderType string // 表名，Index表里定义的类型

	OriginalHeaderType string // HeaderFields对应的ObjectType，KV表为TableField

	FileName string

	SheetName string

	Rows []*DataRow // 0 下标为表头数据

	Headers []*HeaderField
}

// ArrayFieldCount 重复列在表中的索引, 相对于重复列的数量
func (selfObj *DataTable) ArrayFieldCount(field *HeaderField) (ret int) {

	for _, hf := range selfObj.Headers {
		if hf.TypeInfo != nil && hf.TypeInfo.FieldName == field.TypeInfo.FieldName {

			ret++
		}
	}

	return
}

// DataRowIndex 模板用，排除表头的数据索引
func (selfObj *DataTable) DataRowIndex() (ret []int) {

	numRows := len(selfObj.Rows)

	if numRows == 0 {
		return
	}

	ret = make([]int, numRows-1)

	// 排除表头数据
	for i := 0; i < numRows-1; i++ {
		ret[i] = i + 1
	}

	return
}

func (selfObj *DataTable) String() string {

	var sb strings.Builder
	sb.WriteString("====DataTable====\n")
	sb.WriteString(fmt.Sprintf("HeaderType: %s\n", selfObj.HeaderType))
	sb.WriteString(fmt.Sprintf("OriginalHeaderType: %s\n", selfObj.OriginalHeaderType))
	sb.WriteString(fmt.Sprintf("FileName: %s\n", selfObj.FileName))
	sb.WriteString(fmt.Sprintf("SheetName: %s\n", selfObj.SheetName))

	// 遍历所有行
	for row, rowData := range selfObj.Rows {

		sb.WriteString(fmt.Sprintf("%d ", row))

		// 遍历一行中的所有列值
		for index, cell := range rowData.Cells() {

			if index > 0 {
				sb.WriteString("/")
			}

			sb.WriteString(cell.Value)

		}

		sb.WriteString("\n")
	}

	return sb.String()
}

func (selfObj *DataTable) MustGetHeader(col int) (header *HeaderField) {

	for len(selfObj.Headers) <= col {
		selfObj.Headers = append(selfObj.Headers, &HeaderField{
			Cell: &Cell{
				Col: len(selfObj.Headers),
			},
		})
	}

	return selfObj.HeaderByColumn(col)
}

func (selfObj *DataTable) HeaderByColumn(col int) *HeaderField {

	if col >= len(selfObj.Headers) {
		return nil
	}

	return selfObj.Headers[col]
}

func (selfObj *DataTable) HeaderByName(name string) *HeaderField {
	for _, header := range selfObj.Headers {

		if header.TypeInfo == nil {
			continue
		}

		if header.TypeInfo.Name == name || header.TypeInfo.FieldName == name {
			return header
		}
	}

	return nil
}

func (selfObj *DataTable) AddRow() (row int) {

	row = len(selfObj.Rows)

	selfObj.Rows = append(selfObj.Rows, newDataRow(row, selfObj))

	return
}

func (selfObj *DataTable) AddCell(row int) *Cell {

	if row >= len(selfObj.Rows) {
		return nil
	}

	rowData := selfObj.Rows[row]

	return rowData.AddCell()
}

func (selfObj *DataTable) MustGetCell(row, col int) *Cell {

	for len(selfObj.Rows) <= row {
		selfObj.AddRow()
	}

	rowData := selfObj.Rows[row]
	for len(rowData.cells) <= col {
		rowData.AddCell()
	}

	return rowData.Cell(col)
}

// GetCell 代码生成专用
func (selfObj *DataTable) GetCell(row, col int) *Cell {

	if row >= len(selfObj.Rows) {
		return nil
	}

	rowData := selfObj.Rows[row]

	if col >= len(rowData.cells) {
		return nil
	}

	return rowData.Cell(col)
}

// GetValueByName 根据列头找到该行对应的值
func (selfObj *DataTable) GetValueByName(row int, name string) *Cell {

	header := selfObj.HeaderByName(name)

	if header == nil {
		return nil
	}

	return selfObj.GetCell(row, header.Cell.Col)
}

func NewDataTable() *DataTable {
	return &DataTable{}
}
