package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExcelFileLoad(t *testing.T) {

	strCol := ConvertNumToChar(1)
	assert.Equal(t, "A", strCol, "ConvertNumToChar(1) failed")

	strCol = ConvertNumToChar(27)
	assert.Equal(t, "AA", strCol, "ConvertNumToChar(1) failed")

	for i := 1; i < 1000; i++ {
		strCol = ConvertNumToChar(i)

		num, err := ConvertCharToNum(strCol)
		assert.Nilf(t, err, " ConvertCharToNum(%s) failed", strCol)
		assert.Equal(t, i, num, "ConvertCharToNum(%s) failed", strCol)
	}

	_, err := ConvertCharToNum("1.23456789")
	assert.NotNil(t, err, "GeneralNumericScientific failed")

	numValue, err := GeneralNumericScientific("1.23456789", false)
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
