package FlagParam

import (
	"flag"
	"fmt"
)

// 标准参数
var (
	// ParamVersion 显示版本号
	ParamVersion = flag.Bool("version", false, "Show version")

	// ParamExcelResPath 资源路径
	ParamExcelResPath = flag.String("res_path", "", "")
	
	// ParamLanguage 输出日志语言
	ParamLanguage = flag.String("lan", "en_us", "set output language")
)

var (
	// ParamPackageName 导出代码中包/命名空间名称
	ParamPackageName = flag.String("package", "", "override the package name in table @Types")
)

var (
	// Version 版本号
	Version = "0.0.1"
)

func BuildInfoPrint() {
	fmt.Println("Version: ", Version)
}
