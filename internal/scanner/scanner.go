package scanner

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/aqyuki/docat/internal/display"
	"github.com/aqyuki/docat/internal/document"
	"github.com/aqyuki/docat/internal/scanner/options"
	"github.com/aqyuki/goutil/files"
	"github.com/briandowns/spinner"
)

// listAllFile return list of add file in root directory and sub directory
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

// listRootFiles return list of add file in root directory
func listRootFiles(root string) []string {
	list := make([]string, 0)

	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("%+v", err)
		return nil
	}

CHECK:
	for _, file := range files {
		if file.IsDir() {
			continue CHECK
		}

		info, err := file.Info()
		if err != nil {
			continue CHECK
		}

		list = append(list, info.Name())
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
		if document.CONTRIBUTE.Match(item) {
			items = append(items, item)
			continue CHECK_LOOP
		}
	}
	return items
}

// createDocumentList create detected documents in current directory
func createDocumentList(path string, opts *options.ScannerOption) ([]string, error) {
	exist, err := files.ExistDir(path)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, nil
	}

	// Panic may occur if (opts != nil) and (opts.IgnoreSubDirectory) are swapped (panic: runtime error: invalid memory address or nil pointer dereference)
	if opts != nil && opts.IgnoreSubDirectory {
		return extractListItem(listRootFiles(path)), nil
	}

	return extractListItem(listAllFiles(path)), nil
}

// ShowDocumentList show detected documents in current directory
func ShowDocumentList(path string, opts *options.ScannerOption) error {
	files, err := createDocumentList(path, opts)
	if err != nil {
		return err
	}
	display.ShowSimpleList(files)
	return nil
}

// ScanWithSpinner scan selected directory with spinner
func ScanWithSpinner(path string, opts *options.ScannerOption) ([]string, error) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	data, err := createDocumentList(path, opts)
	s.Stop()
	return data, err
}
