package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockTableSheet struct {
	name    string
	maxCol  int
	maxRows int
	values  map[int]map[int]string
}

func (m *MockTableSheet) Name() string {
	return m.name
}

func (m *MockTableSheet) GetValue(row, col int, opt *ValueOption) string {
	if m.values == nil {
		return ""
	}
	if rowValues, ok := m.values[row]; ok {
		if value, ok := rowValues[col]; ok {
			return value
		}
	}
	return ""
}

func (m *MockTableSheet) MaxColumn() int {
	return m.maxCol
}

func (m *MockTableSheet) IsRowEmpty(row, maxCol int) bool {
	if maxCol == -1 {
		maxCol = m.maxCol
	}
	for col := 0; col < maxCol; col++ {
		if m.GetValue(row, col, nil) != "" {
			return false
		}
	}
	return true
}

func TestReadSheetRow(t *testing.T) {
	sheet := &MockTableSheet{
		maxCol: 3,
		values: map[int]map[int]string{
			0: {0: "A1", 1: "B1", 2: "C1"},
			1: {0: "A2", 1: "B2", 2: "C2"},
		},
	}

	row := ReadSheetRow(sheet, 0)
	assert.Equal(t, []string{"A1", "B1", "C1"}, row, "ReadSheetRow failed")

	row = ReadSheetRow(sheet, 1)
	assert.Equal(t, []string{"A2", "B2", "C2"}, row, "ReadSheetRow failed")

	row = ReadSheetRow(sheet, 2)
	assert.Equal(t, []string{"", "", ""}, row, "ReadSheetRow failed")
}
