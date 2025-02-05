package report

import (
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
