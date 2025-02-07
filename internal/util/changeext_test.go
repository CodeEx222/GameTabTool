package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeExtension(t *testing.T) {
	result := ChangeExtension("example.txt", ".md")
	assert.Equal(t, "example.md", result, "ChangeExtension failed")

	result = ChangeExtension("example", ".md")
	assert.Equal(t, "example.md", result, "ChangeExtension failed")

	result = ChangeExtension("example.txt", "")
	assert.Equal(t, "example", result, "ChangeExtension failed")
}

func TestConvertNumToChar(t *testing.T) {

	strCol := ConvertNumToChar(1)
	assert.Equal(t, "A", strCol, "ConvertNumToChar(1) failed")

	strCol = ConvertNumToChar(27)
	assert.Equal(t, "AA", strCol, "ConvertNumToChar(27) failed")

}

func TestConvertCharToNum(t *testing.T) {

	for i := 1; i < 1000; i++ {
		strCol := ConvertNumToChar(i)
		num, err := ConvertCharToNum(strCol)
		assert.Nilf(t, err, "ConvertCharToNum(%s) failed", strCol)
		assert.Equal(t, i, num, "ConvertCharToNum(%s) failed", strCol)
	}

	_, err := ConvertCharToNum("1.23456789")
	assert.NotNil(t, err, "ConvertCharToNum with invalid input failed")

}

func TestConvertExcelCellToNumPos(t *testing.T) {
	result, xCol, yRow := ConvertExcelCellToNumPos("")
	assert.False(t, result, "ConvertExcelCellToNumPos failed")
	assert.Equal(t, 0, xCol, "ConvertExcelCellToNumPos failed")
	assert.Equal(t, 0, yRow, "ConvertExcelCellToNumPos failed")

	result, xCol, yRow = ConvertExcelCellToNumPos("A1")
	assert.True(t, result, "ConvertExcelCellToNumPos failed")
	assert.Equal(t, 1, xCol, "ConvertExcelCellToNumPos failed")
	assert.Equal(t, 1, yRow, "ConvertExcelCellToNumPos failed")

	result, xCol, yRow = ConvertExcelCellToNumPos("B2")
	assert.True(t, result, "ConvertExcelCellToNumPos failed")
	assert.Equal(t, 2, xCol, "ConvertExcelCellToNumPos failed")
	assert.Equal(t, 2, yRow, "ConvertExcelCellToNumPos failed")

	result, xCol, yRow = ConvertExcelCellToNumPos("@3")
	assert.False(t, result, "ConvertExcelCellToNumPos failed")
	assert.Equal(t, 0, xCol, "ConvertExcelCellToNumPos failed")
	assert.Equal(t, 0, yRow, "ConvertExcelCellToNumPos failed")

	result, xCol, yRow = ConvertExcelCellToNumPos("A@")
	assert.False(t, result, "ConvertExcelCellToNumPos failed")
	assert.Equal(t, 0, xCol, "ConvertExcelCellToNumPos failed")
	assert.Equal(t, 0, yRow, "ConvertExcelCellToNumPos failed")

}

func TestGeneralNumericScientific(t *testing.T) {

	numValue, err := GeneralNumericScientific("", true)
	assert.Nil(t, err, "GeneralNumericScientific failed")
	assert.Equal(t, "", numValue, "GeneralNumericScientific failed")

	numValue, err = GeneralNumericScientific("a2.34", true)
	assert.NotNil(t, err, "GeneralNumericScientific failed")
	assert.Equal(t, "a2.34", numValue, "GeneralNumericScientific failed")

	numValue, err = GeneralNumericScientific("1.23456789", false)
	assert.Nil(t, err, "GeneralNumericScientific failed")
	assert.Equal(t, "1.23456789", numValue, "GeneralNumericScientific failed")

	numValue, err = GeneralNumericScientific("123456789", false)
	assert.Nil(t, err, "GeneralNumericScientific failed")
	assert.Equal(t, "123456789", numValue, "GeneralNumericScientific failed")

	numValue, err = GeneralNumericScientific("-123456789", false)
	assert.Nil(t, err, "GeneralNumericScientific failed")
	assert.Equal(t, "-123456789", numValue, "GeneralNumericScientific failed")

	numValue, err = GeneralNumericScientific("1.23456789", true)
	assert.Nil(t, err, "GeneralNumericScientific failed")
	assert.Equal(t, "1.23456789", numValue, "GeneralNumericScientific failed")

	numValue, err = GeneralNumericScientific("0.000000000134", true)
	assert.Nil(t, err, "GeneralNumericScientific failed")
	assert.Equal(t, "1.34E-10", numValue, "GeneralNumericScientific failed")

	numValue, err = GeneralNumericScientific("120000000000", true)
	assert.Nil(t, err, "GeneralNumericScientific failed")
	assert.Equal(t, "1.2E+11", numValue, "GeneralNumericScientific failed")

}
