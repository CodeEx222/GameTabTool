package report

type ErrorLanguage struct {
	CHS string
	EN  string
}

type ErrorID string

var (
	UnknownInputFileExtension ErrorID = "UnknownInputFileExtension"
	UnknownSheet              ErrorID = "UnknownSheet"
)

var (
	ErrorByID = map[ErrorID]*ErrorLanguage{
		UnknownInputFileExtension: {CHS: "未知的输入文件扩展名", EN: "Unknown input file extension"},
		UnknownSheet:              {CHS: "未找到 [Sheet1] 表 请检查文件", EN: "The [Sheet1] table not found Please check the file"},
	}
)
