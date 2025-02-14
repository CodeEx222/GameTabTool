package FlagParam

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// langPacks 定义语言包  key - [英文, 中文]
var langPacks = map[string][]string{
	"usage":    {"%s \nUsage of:\n", "%s \n用法:\n"},
	"version":  {"Show version", "显示版本号"},
	"res_path": {"Set Resource path", "设置资源路径"},
	"lan":      {"Language preference (en, zh)", "语言首选项(en, zh)"},
	"index":    {"Set index file name", "设置索引文件名"},
}

// 标准参数
var (
	// ParamVersion 显示版本号
	ParamVersion bool

	// ParamExcelResPath 资源路径
	ParamExcelResPath string

	// ParamLanguage 输出日志语言
	ParamLanguage string
)

var (
	// ParamIndexName 索引文件名
	ParamIndexName string
)

func Parse() {

	// 提前解析 -lang 参数
	for _, arg := range os.Args[1:] {
		// 检查参数是否以 -lan 或 --lan 开头
		if strings.HasPrefix(arg, "-lan=") || strings.HasPrefix(arg, "--lan=") {
			// 提取 = 后面的值
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) == 2 {
				ParamLanguage = parts[1]
			}
		}
	}

	// 检查语言是否支持
	langIndex := 0
	if ParamLanguage != "en" && ParamLanguage != "zh" {
		fmt.Printf("Unsupported language: %s, defaulting to 'en'\n", ParamLanguage)
		ParamLanguage = "en"
	}

	// 根据语言包设置参数描述
	if ParamLanguage == "zh" {
		langIndex = 1
	}

	// 自定义 Usage 函数
	flag.Usage = func() {
		_, err := fmt.Fprintf(os.Stderr, langPacks["usage"][langIndex], os.Args[0])
		if err != nil {
			return
		}
		flag.PrintDefaults()
	}

	// 根据语言包设置参数描述
	flag.StringVar(&ParamExcelResPath, "res_path", "", langPacks["res_path"][langIndex])
	flag.StringVar(&ParamLanguage, "lan", "en", langPacks["lan"][langIndex])
	flag.BoolVar(&ParamVersion, "version", false, langPacks["version"][langIndex])

	flag.StringVar(&ParamIndexName, "index", "", langPacks["index"][langIndex])

	// 解析命令行参数
	flag.Parse()
}

var (
	// Version 版本号
	Version = "0.0.1"
)

func BuildInfoPrint() {
	fmt.Println("Version: ", Version)
}
