package model

type IndexDefine struct {
	TableType     string   // 表类型
	TableFileName string   // 表文件名
	Group         []string // 标记 | 分割
	Comment       string   // 注释
}

func (selfObj *IndexDefine) ContainTag(tag string) bool {
	for _, s := range selfObj.Group {
		if s == tag {
			return true
		}
	}
	return false
}
