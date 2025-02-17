package report

import (
	"fmt"
	"github.com/charmbracelet/log"
	"os"
	"time"
)

var ToolsLog *log.Logger

func init() {
	ToolsLog = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.DateTime,
	})
	// 设置日志级别
	log.SetLevel(log.DebugLevel)
}

func LogFatal(msg interface{}, keyVals ...interface{}) {
	ToolsLog.Log(log.FatalLevel, msg, keyVals...)
	panic("LogFatal called")
}

func LogFatalf(format string, args ...interface{}) {
	ToolsLog.Log(log.FatalLevel, fmt.Sprintf(format, args...))
	panic("LogFatalf called")

}
