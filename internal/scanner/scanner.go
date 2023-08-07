package scanner

import (
	"fmt"
	"io/fs"
	"path/filepath"
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

func ScanDocument(path string) error {
	return nil
}
