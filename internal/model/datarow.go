package model

type DataRow struct {
	row   int
	cells []*Cell
	tab   *DataTable
}

func (selfObj *DataRow) Cells() []*Cell {
	return selfObj.cells
}

func (selfObj *DataRow) Cell(col int) *Cell {
	return selfObj.cells[col]
}

func (selfObj *DataRow) AddCell() (ret *Cell) {

	ret = &Cell{
		Col:   len(selfObj.cells),
		Row:   selfObj.row,
		Table: selfObj.tab,
	}

	selfObj.cells = append(selfObj.cells, ret)
	return
}

func (selfObj *DataRow) IsEmpty() bool {
	return len(selfObj.cells) == 0
}

func newDataRow(row int, tab *DataTable) *DataRow {
	return &DataRow{
		row: row,
		tab: tab,
	}
}
