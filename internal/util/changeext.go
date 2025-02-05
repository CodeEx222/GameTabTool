package util

import (
	"errors"
	"math"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func ChangeExtension(filename, newExt string) string {

	file := filepath.Base(filename)

	return strings.TrimSuffix(file, path.Ext(file)) + newExt
}

var ExcelChar = []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// ConvertNumToChar 将数字转换为 Excel 格式的列名。
func ConvertNumToChar(num int) string {
	if num < 27 {
		return ExcelChar[num]
	}
	k := num % 26
	if k == 0 {
		k = 26
	}
	v := (num - k) / 26
	col := ConvertNumToChar(v)

	cols := col + ExcelChar[k]
	return cols
}

func ConvertCharToNum(col string) (int, error) {
	if len(col) == 0 {
		return 0, nil
	}

	// 26 进制转 10 进制
	num := 0
	for i := 0; i < len(col); i++ {
		// 检测字符串是否在 A -  Z 之间
		if col[i] < 'A' || col[i] > 'Z' {
			return 0, errors.New("not found")
		}

		// ExcelChar 里的第一个元素是空字符串，所以要 +1
		// ExcelChar 找到 col[i] 的索引，就是 col[i] 对应的数字
		numIndex := 0
		for j := 1; j < len(ExcelChar); j++ {
			if ExcelChar[j] == string(col[i]) {
				numIndex = j
				break
			}
		}

		num = num*26 + numIndex
	}
	return num, nil
}

const (
	maxNonScientificNumber = 1e11
	minNonScientificNumber = 1e-9
)

// GeneralNumericScientific 将数字转换为科学计数法或普通数字。
func GeneralNumericScientific(value string, allowScientific bool) (string, error) {
	if strings.TrimSpace(value) == "" {
		return "", nil
	}
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return value, err
	}
	if allowScientific {
		absF := math.Abs(f)
		// When using General format, numbers that are less than 1e-9 (0.000000001) and greater than or equal to
		// 1e11 (100,000,000,000) should be shown in scientific notation.
		// Numbers less than the number after zero, are assumed to be zero.
		if (absF >= math.SmallestNonzeroFloat64 && absF < minNonScientificNumber) || absF >= maxNonScientificNumber {
			return strconv.FormatFloat(f, 'E', -1, 64), nil
		}
	}
	// This format (fmt="f", prec=-1) will prevent padding with zeros and will never switch to scientific notation.
	// However, it will show more than 11 characters for very precise numbers, and this cannot be changed.
	// You could also use fmt="g", prec=11, which doesn't pad with zeros and allows the correct precision,
	// but it will use scientific notation on numbers less than 1e-4. That value is hardcoded in Go and cannot be
	// configured or disabled.
	return strconv.FormatFloat(f, 'f', -1, 64), nil
}
