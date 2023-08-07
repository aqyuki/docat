package scanner_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/aqyuki/docat/internal/scanner"
	"github.com/aqyuki/goutil/files"
)

func Test_listAllFiles(t *testing.T) {
	const (
		fileCount = 4
	)
	_, currentFile, _, _ := runtime.Caller(0)
	t.Logf("%+v\n", currentFile)
	testDir := filepath.Join(filepath.Dir(currentFile), "testdata")

	list := scanner.ListAllFiles(testDir)

	count := len(list)
	if count != fileCount {
		t.Errorf("The number of files read is abnormal. want -> %d, got -> %d\n", fileCount, count)
	}

	for _, path := range list {
		exist, err := files.ExistFile(path)
		if err != nil {
			t.Errorf("An unexpected error has occurred. %+v\n", err)
		}
		if !exist {
			t.Errorf("A non-existent file was returned. got -> %+v\n", path)
		}
	}
}
