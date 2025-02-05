package helper

type TableFile interface {
	Load(filename string) error

	// Save 保存到文件
	Save(filename string) error

	// Sheets 获取所有表单
	Sheets() []TableSheet
}

type ValueOption struct {
	ValueAsFloat bool
}

type TableSheet interface {

	// Name 表单名称
	Name() string

	// GetValue 从表单指定单元格获取值
	GetValue(row, col int, opt *ValueOption) string

	// MaxColumn 最大列
	MaxColumn() int

	// IsRowEmpty 检测本行是否全空(结束)
	IsRowEmpty(row, maxCol int) bool
}

func ReadSheetRow(sheet TableSheet, row int) (ret []string) {

	ret = make([]string, sheet.MaxColumn())
	for col := 0; col < sheet.MaxColumn(); col++ {

		value := sheet.GetValue(row, col, nil)

		ret[col] = value
	}

	return
}
