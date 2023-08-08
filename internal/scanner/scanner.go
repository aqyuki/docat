package scanner

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/aqyuki/docat/internal/document"
	"github.com/aqyuki/docat/internal/printer"
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

// extractListItem extracts only the documents to be displayed
func extractListItem(list []string) []string {
	items := make([]string, 0)
CHECK_LOOP:
	for _, item := range list {
		if document.README.Match(item) {
			items = append(items, item)
			continue CHECK_LOOP
		}

		if document.LICENSE.Match(item) {
			items = append(items, item)
			continue CHECK_LOOP
		}

		if document.CHANGELOG.Match(item) {
			items = append(items, item)
			continue CHECK_LOOP
		}
		if document.CONTRIBUTING.Match(item) {
			items = append(items, item)
			continue CHECK_LOOP
		}
		if document.CONTRIBUTOR.Match(item) {
			items = append(items, item)
			continue CHECK_LOOP
		}
	}
	return items
}

func ShowDocumentList(path string) error {
	exist, err := files.ExistDir(path)
	if err != nil {
		return nil
	}
	if !exist {
		return nil
	}
	printer.PrintListNonSelectable(extractListItem(listAllFiles(path)))
	return nil
}
