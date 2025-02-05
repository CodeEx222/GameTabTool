package main

import (
	"flag"
	"gametabtool/internal/FlagParam"
	"gametabtool/internal/entry"
)

func main() {

	// 处理所有输入参数
	flag.Parse()

	// 版本
	if *FlagParam.ParamVersion {
		FlagParam.BuildInfoPrint()
		return
	}

	// 开始处理
	entry.Entry()

}
