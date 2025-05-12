package entry

import (
	"gametabtool/internal/FlagParam"
	"gametabtool/internal/compiler"
	"gametabtool/internal/helper"
	"gametabtool/internal/model"
	"gametabtool/internal/report"
	"os"
	"path"
	"strings"
)

func Entry() {

	projectDir, _ := os.Getwd()

	projectDir = strings.Replace(projectDir, "\\", "/", -1)

	// 初始化全局数据
	model.GlobalData = model.NewGlobals()
	model.GlobalData.ExcelDataPath = FlagParam.ParamExcelResPath
	model.GlobalData.IndexDataPath = FlagParam.ParamIndexName

	model.GlobalData.ExcelDataPath = path.Join(projectDir, model.GlobalData.ExcelDataPath)
	model.GlobalData.ExcelDataPath = path.Clean(model.GlobalData.ExcelDataPath)

	// 读取索引文件
	idxloader := helper.NewFileLoader(true)
	model.GlobalData.IndexGetter = idxloader

	var err error
	err = compiler.Compile()

	if err != nil {
		goto Exit
	}

	//err = check.RunCheckTableDataCheck(globals)

	//if err != nil {
	//	goto Exit
	//}

	// 写入文件

	return
Exit:
	report.ToolsLog.Error(err)
	os.Exit(1)
}
