package helper

import (
	"errors"
	_ "gametabtool/test_init"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type MockTableFile struct{}

func (m *MockTableFile) Load(filename string) error {
	return nil
}

func (m *MockTableFile) Save(filename string) error {
	return nil
}

func (m *MockTableFile) Sheets() []TableSheet {
	return nil
}

func NewMockTableFile() TableFile {
	return &MockTableFile{}
}

func TestFileLoader_AddFile(t *testing.T) {
	loader := NewFileLoader(false)
	loader.AddFile("test.xlsx")
	assert.Equal(t, 1, len(loader.inputFile), "File not added correctly")
}

func TestFileLoader_Commit(t *testing.T) {
	loader := NewFileLoader(false)
	loader.AddFile("test.xlsx")
	loader.Commit()
	_, ok := loader.fileByName.Load("test.xlsx")
	assert.True(t, ok, "File not loaded correctly")
}

func TestFileLoader_GetFile_SyncLoad(t *testing.T) {

	projectDir, _ := os.Getwd()

	loader := NewFileLoader(true)
	file, err := loader.GetFile(projectDir + "/bin/testdata/testRead.xlsx")
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, file, "File should not be nil")
}

func TestFileLoader_GetFile_AsyncLoad(t *testing.T) {
	projectDir, _ := os.Getwd()

	loader := NewFileLoader(false)
	loader.AddFile(projectDir + "/bin/testdata/testRead.xlsx")
	loader.Commit()
	file, err := loader.GetFile(projectDir + "/bin/testdata/testRead.xlsx")
	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, file, "File should not be nil")
}

func TestFileLoader_GetFile_NotFound(t *testing.T) {
	loader := NewFileLoader(false)
	_, err := loader.GetFile("nonexistent.xlsx")
	assert.NotNil(t, err, "Error should not be nil")
	assert.Equal(t, errors.New("not found"), err, "Error message should be 'not found'")
}

func TestLoadFileByExt(t *testing.T) {
	projectDir, _ := os.Getwd()
	result := loadFileByExt(projectDir + "/bin/testdata/testRead.xlsx")
	assert.NotNil(t, result, "Result should not be nil")
	_, ok := result.(TableFile)
	assert.True(t, ok, "Result should be of type TableFile")

	assert.Fail(t, "Need to test the case when the file extension is not supported")
}

func TestNewFileLoader(t *testing.T) {
	loader := NewFileLoader(true)
	assert.NotNil(t, loader, "Loader should not be nil")
	assert.True(t, loader.syncLoad, "syncLoad should be true")
}
