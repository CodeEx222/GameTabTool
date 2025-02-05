package helper

import (
	"errors"
	"gametabtool/internal/report"
	"path/filepath"
	"sync"
)

type FileGetter interface {
	GetFile(filename string) (TableFile, error)
}

type FileLoader struct {
	fileByName sync.Map
	inputFile  []string

	syncLoad bool
	cacheDir string
}

func (selfObj *FileLoader) AddFile(filename string) {

	selfObj.inputFile = append(selfObj.inputFile, filename)
}

func (selfObj *FileLoader) Commit() {

	var task sync.WaitGroup
	task.Add(len(selfObj.inputFile))

	for _, inputFileName := range selfObj.inputFile {

		go func(fileName string) {

			selfObj.fileByName.Store(fileName, loadFileByExt(fileName, selfObj.cacheDir))

			task.Done()

		}(inputFileName)

	}

	task.Wait()

	selfObj.inputFile = selfObj.inputFile[0:0]
}

func loadFileByExt(filename string, cacheDir string) interface{} {

	var tabFile TableFile
	switch filepath.Ext(filename) {
	case ".xlsx", ".xls", ".xlsm":

		tabFile = NewXlsxFile(cacheDir)

		err := tabFile.Load(filename)

		if err != nil {
			return err
		}

	default:
		report.LogTableError(report.UnknownInputFileExtension, filename)
	}

	return tabFile
}

func (selfObj *FileLoader) GetFile(filename string) (TableFile, error) {

	if selfObj.syncLoad {

		result := loadFileByExt(filename, selfObj.cacheDir)
		if err, ok := result.(error); ok {
			return nil, err
		}

		return result.(TableFile), nil

	} else {
		if result, ok := selfObj.fileByName.Load(filename); ok {

			if err, ok := result.(error); ok {
				return nil, err
			}

			return result.(TableFile), nil

		} else {
			return nil, errors.New("not found")
		}
	}

}

func NewFileLoader(syncLoad bool, cacheDir string) *FileLoader {
	return &FileLoader{
		syncLoad: syncLoad,
		cacheDir: cacheDir,
	}
}
