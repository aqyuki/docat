package scanner

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/aqyuki/goutil/files"
)

// listAllFile return list of add file in root directory
func listAllFiles(root string) []string {
	list := make([]string, 0)

	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			list = append(list, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("%+v", err)
		return nil
	}

	return list
}

func ShowDocumentList(path string) error {
	exist,err := files.ExistDir(path)

	if err != nil{
		return nil
	}
	if !exist{
		return nil
	}

	list := listAllFiles(path)
	if len(list) < 1{
		return nil
	}
	return nil
}
