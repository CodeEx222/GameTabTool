package helper

import (
	_ "gametabtool/test_init"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExcelFileLoad(t *testing.T) {
	projectDir, _ := os.Getwd()

	var excelFile XlsxFile
	err := excelFile.Load(projectDir + "/bin/testdata/testRead.xlsx")
	assert.Nil(t, err, "读取文件失败")

	// test read
	sheet := excelFile.Sheets()
	assert.Nil(t, err, "excel 里没有sheet")

	assert.EqualValues(t, 2, len(sheet), "sheet 数量不对")
	assert.EqualValues(t, "Sheet1", sheet[0].Name(), "sheet1 名字不对")
	assert.EqualValues(t, "Sheet2", sheet[1].Name(), "sheet2 名字不对")

	assert.Equal(t, 3, sheet[0].MaxColumn(), "sheet1 列数不对")
	assert.Equal(t, 8, sheet[1].MaxColumn(), "sheet2 列数不对")
}
