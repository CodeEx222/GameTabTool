package main

import (
	"gametabtool/internal/FlagParam"
	"gametabtool/internal/entry"
)

func main() {

	// 处理所有输入参数
	FlagParam.Parse()

	// 版本
	if FlagParam.ParamVersion {
		FlagParam.BuildInfoPrint()
		return
	}

	// 开始处理
	entry.Entry()

}
