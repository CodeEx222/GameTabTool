package helper

import (
	_ "gametabtool/test_init"
	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
	"os"
	"testing"
)

func TestXlsxFile_Load(t *testing.T) {
	projectDir, _ := os.Getwd()

	var xlsxFile XlsxFile
	err := xlsxFile.Load(projectDir + "/bin/testdata/testRead.xlsx")
	assert.Nil(t, err, "Failed to load file")

	sheets := xlsxFile.Sheets()
	assert.Equal(t, 2, len(sheets), "Incorrect number of sheets")
	assert.Equal(t, "Sheet1", sheets[0].Name(), "Incorrect sheet name")
	assert.Equal(t, "Sheet2", sheets[1].Name(), "Incorrect sheet name")

	assert.Equal(t, 3, sheets[0].MaxColumn(), "Incorrect sheet MaxColumn")
	assert.Equal(t, 8, sheets[1].MaxColumn(), "Incorrect sheet MaxColumn")

}

func TestXlsxFile_Save(t *testing.T) {
	projectDir, _ := os.Getwd()

	var xlsxFile XlsxFile
	err := xlsxFile.Load(projectDir + "/bin/testdata/testRead.xlsx")
	assert.Nil(t, err, "Failed to load file")

	err = xlsxFile.Save(projectDir + "/bin/testdata/testSave.xlsx")
	assert.Nil(t, err, "Failed to save file")

	_, err = os.Stat(projectDir + "/bin/testdata/testSave.xlsx")
	assert.False(t, os.IsNotExist(err), "File was not saved")
}

func TestXlsxFile_FromXFile(t *testing.T) {
	projectDir, _ := os.Getwd()

	file, err := excelize.OpenFile(projectDir + "/bin/testdata/testRead.xlsx")
	assert.Nil(t, err, "Failed to open file")

	var xlsxFile XlsxFile
	xlsxFile.FromXFile(file)

	sheets := xlsxFile.Sheets()
	assert.Equal(t, 2, len(sheets), "Incorrect number of sheets")
	assert.Equal(t, "Sheet1", sheets[0].Name(), "Incorrect sheet name")
	assert.Equal(t, "Sheet2", sheets[1].Name(), "Incorrect sheet name")
}

func TestXlsxSheet_GetValue(t *testing.T) {
	projectDir, _ := os.Getwd()

	file, err := excelize.OpenFile(projectDir + "/bin/testdata/testRead.xlsx")
	assert.Nil(t, err, "Failed to open file")

	sheet, _ := newXlsxSheet("Sheet1", file)
	value := sheet.GetValue(0, 0, nil)
	assert.Equal(t, "种类", value, "Incorrect cell value")

	valueNum := sheet.GetValue(1, 0, &ValueOption{ValueAsFloat: true})
	assert.Equal(t, "10.2", valueNum, "Incorrect cell value")

	assert.Panicsf(t, func() {
		newXlsxSheet("Sheet1113", file)
	}, "Should panic")

	_, sheet3Result := newXlsxSheet("Sheet3", file)
	assert.False(t, sheet3Result, "Sheet should be null")

	assert.Panicsf(t, func() {
		sheet.GetValue(-1, 0, nil)
	}, "Should panic")

}

func TestXlsxSheet_IsRowEmpty(t *testing.T) {
	projectDir, _ := os.Getwd()

	file, err := excelize.OpenFile(projectDir + "/bin/testdata/testRead.xlsx")
	assert.Nil(t, err, "Failed to open file")

	sheet, _ := newXlsxSheet("Sheet1", file)
	isEmpty := sheet.IsRowEmpty(0, -1)
	assert.False(t, isEmpty, "Row should not be empty")

	isEmpty = sheet.IsRowEmpty(10, -1)
	assert.True(t, isEmpty, "Row should be empty")
}
