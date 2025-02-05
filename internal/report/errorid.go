package report

type ErrorLanguage struct {
	CHS string
}

type ErrorID string

var (
	UnknownInputFileExtension ErrorID = "UnknownInputFileExtension"
)

var (
	ErrorByID = map[ErrorID]*ErrorLanguage{
		UnknownInputFileExtension: {CHS: "未知的输入文件扩展名"},
	}
)
