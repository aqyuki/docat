package scanner_test

import (
	"path/filepath"
	"reflect"
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

func Test_extractListItem(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "All match",
			args: args{
				list: []string{
					filepath.Join("project", "README"),
					filepath.Join("project", "LICENSE"),
					filepath.Join("project", "CHANGELOG"),
					filepath.Join("project", "CONTRIBUTING"),
					filepath.Join("project", "CONTRIBUTOR"),
					filepath.Join("project", "CONTRIBUTE"),
				},
			},
			want: []string{
				filepath.Join("project", "README"),
				filepath.Join("project", "LICENSE"),
				filepath.Join("project", "CHANGELOG"),
				filepath.Join("project", "CONTRIBUTING"),
				filepath.Join("project", "CONTRIBUTOR"),
				filepath.Join("project", "CONTRIBUTE"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scanner.ExtractListItem(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractListItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
