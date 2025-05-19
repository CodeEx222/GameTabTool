package compiler

import (
	"gametabtool/internal/model"
	"gametabtool/internal/report"
	"sort"
)

func Compile() (ret error) {

	defer func() {

		switch err := recover().(type) {
		case *report.TableError:
			ret = err
		case nil:
		default:
			panic(err)
		}

	}()

	// 注册默认格式
	model.InitBuiltinTypes()

	report.ToolsLog.Debugf("Loading Index file: '%s'... ", model.GlobalData.IndexDataPath)
	err := LoadIndexTable(model.GlobalData)

	if err != nil {
		return err
	}

	// 排序 model.GlobalData.IndexDefine 优先把TableType 为空的放在前面
	// 这样可以保证在加载表的时候, 先加载没有依赖的表

	// 排序 model.GlobalData.IndexDefine 优先把 TableType 为空的放在前面
	sort.Slice(model.GlobalData.IndexDefine, func(i, j int) bool {
		return model.GlobalData.IndexDefine[i].TableType == "" && model.GlobalData.IndexDefine[j].TableType != ""
	})

	// 读取所有要导出的文件, 把所有结构体都加载到内存中

	for _, indexAll := range model.GlobalData.IndexDefine {
		if indexAll == nil {
			continue
		}

		if indexAll.TableFileName == "" {
			continue
		}

		report.ToolsLog.Debugf("Loading table file: '%s'... ", indexAll.TableFileName)

		//tabLoader := helper.NewFileLoader(!model.GlobalData.ParaLoading, model.GlobalData.CacheDir)
		//tabLoader.AddFile(indexAll.TableFileName)
		//tabLoader.Commit()
		//
		//model.GlobalData.TableGetter = tabLoader
		//
		//err = LoadTable(model.GlobalData, indexAll)
		//
		//if err != nil {
		//	return err
		//}
	}

	return nil
}
