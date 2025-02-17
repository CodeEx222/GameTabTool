package report

import (
	"fmt"
	"gametabtool/internal/FlagParam"
	"strings"
)

type TableError struct {
	ID ErrorID

	context []interface{}
}

func getErrorDesc(id ErrorID) string {

	if lan, ok := ErrorByID[id]; ok {
		if FlagParam.LanguageIndex == 0 {
			return lan.EN
		} else if FlagParam.LanguageIndex == 1 {
			return lan.CHS
		}
	}

	return ""
}

func (SelfObj *TableError) Error() string {

	var sb strings.Builder

	if FlagParam.LanguageIndex == 0 {
		sb.WriteString("TableError.")
	} else if FlagParam.LanguageIndex == 1 {
		sb.WriteString("表错误.")
	}

	sb.WriteString(string(SelfObj.ID))
	sb.WriteString(" ")
	sb.WriteString(getErrorDesc(SelfObj.ID))

	if len(SelfObj.context) > 0 {
		sb.WriteString(" | ")
	}

	for index, c := range SelfObj.context {
		if index > 0 {
			sb.WriteString(" ")
		}

		sb.WriteString(fmt.Sprintf("%+v", c))
	}

	return sb.String()
}

func LogTableError(id ErrorID, context ...interface{}) {

	LogFatal(&TableError{
		ID:      id,
		context: context,
	})
}
