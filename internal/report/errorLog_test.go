package report

import (
	"bytes"
	"gametabtool/internal/FlagParam"
	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestErrorLog(t *testing.T) {

	var TestError ErrorID
	TestError = "TestError"
	ErrorByID["TestError"] = &ErrorLanguage{CHS: "测试错误 ", EN: "Test error "}

	var buf bytes.Buffer
	ToolsLog.SetOutput(&buf)
	defer func() {
		// 恢复默认的输出
		ToolsLog.SetOutput(os.Stderr)
	}()

	oldLanguageIndex := FlagParam.LanguageIndex
	FlagParam.LanguageIndex = 1
	errorStr1 := "context1"
	errorStr2 := "context2"

	ToolsLog.Log(log.ErrorLevel, &TableError{
		ID:      TestError,
		context: []interface{}{errorStr1, errorStr2},
	})

	// 获取捕获的日志输出
	output := buf.String()
	// 使用 testify 的 assert 包进行断言
	assert.Contains(t, output, "表错误.TestError 测试错误  | context1 context2",
		"日志输出不符合预期")

	FlagParam.LanguageIndex = 0

	buf.Reset()
	ToolsLog.Log(log.ErrorLevel, &TableError{
		ID:      TestError,
		context: []interface{}{errorStr1, errorStr2},
	})

	// 获取捕获的日志输出
	output = buf.String()
	// 使用 testify 的 assert 包进行断言
	assert.Contains(t, output, "TableError.TestError Test error  | context1 context2",
		"日志输出不符合预期")

	buf.Reset()

	ToolsLog.Log(log.ErrorLevel, &TableError{
		ID:      "NotExists",
		context: []interface{}{errorStr1, errorStr2},
	})
	output = buf.String()
	assert.Contains(t, output, "TableError.NotExists  | context1 context2",
		"日志输出不符合预期")

	FlagParam.LanguageIndex = oldLanguageIndex

	assert.Panics(t, func() {
		LogFatalf("Test %s", "error")
	}, "")

}
