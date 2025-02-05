package model

import "gametabtool/internal/helper"

type Globals struct {
	// ExcelDataPath excel文件路径
	ExcelDataPath string

	IndexGetter helper.FileGetter // 索引文件获取器
	TableGetter helper.FileGetter // 其他文件获取器
}

func NewGlobals() *Globals {
	return &Globals{}
}

// GlobalData 声明一个全局变量
var GlobalData *Globals
