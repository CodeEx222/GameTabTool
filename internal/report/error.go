package report

import (
	"fmt"
	"strings"
)

type TableError struct {
	ID ErrorID

	context []interface{}
}

func getErrorDesc(id ErrorID) string {

	if lan, ok := ErrorByID[id]; ok {
		return lan.CHS
	}

	return ""
}

func (SelfObj *TableError) Error() string {

	var sb strings.Builder

	sb.WriteString("TableError.")
	sb.WriteString(string(SelfObj.ID))
	sb.WriteString(" ")
	sb.WriteString(getErrorDesc(SelfObj.ID))
	sb.WriteString(" | ")

	for index, c := range SelfObj.context {
		if index > 0 {
			sb.WriteString(" ")
		}

		sb.WriteString(fmt.Sprintf("%+v", c))
	}

	return sb.String()
}

func LogTableError(id ErrorID, context ...interface{}) {

	ToolsLog.Fatal(&TableError{
		ID:      id,
		context: context,
	})
}
