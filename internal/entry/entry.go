package entry

import (
	"gametabtool/internal/FlagParam"
	"gametabtool/internal/model"
)

func Entry() {

	// 初始化全局数据
	model.GlobalData = model.NewGlobals()
	model.GlobalData.ExcelDataPath = *FlagParam.ParamExcelResPath

}
