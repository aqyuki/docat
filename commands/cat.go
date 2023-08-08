package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aqyuki/docat/internal/document"
	"github.com/aqyuki/docat/internal/loader"
	"github.com/aqyuki/docat/internal/scanner"
	"github.com/aqyuki/docat/internal/tags"
	"github.com/aqyuki/docat/internal/view"
	"github.com/spf13/cobra"
)

var catCommand = &cobra.Command{
	Use:   "cat",
	Short: "Displays the contents of the selected file.",
	Long:  "Displays the contents of the specified file. The available document formats are README, LICENSE, CHANGELOG, CONTRIBUTING, or CONTRIBUTOR. If there are multiple files of the same document format, a separate selector will be used to open a new file.",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) > 1 {
			return fmt.Errorf("too many args")
		}

		targetDir, err := os.Getwd()
		if err != nil {
			return err
		}
		if len(args) == 1 {
			dir, err := filepath.Abs(args[0])
			if err != nil {
				return err
			}
			targetDir = dir
		}

		tag, err := view.RunDocumentSelector()
		if err != nil {
			return err
		}

		files, err := scanner.CreateDocumentList(targetDir)
		if err != nil {
			return err
		}

		var pattern *document.DocumentPat

		switch tag {
		case tags.README:
			pattern = document.README
		case tags.LICENSE:
			pattern = document.LICENSE
		case tags.CHANGELOG:
			pattern = document.CHANGELOG
		case tags.CONTRIBUTING:
			pattern = document.CONTRIBUTING
		case tags.CONTRIBUTOR:
			pattern = document.CONTRIBUTOR
		case tags.NON:
			return nil
		}

		docs := make([]string, 0)
		for _, file := range files {
			if pattern.Match(file) {
				docs = append(docs, file)
			}
		}

		var path string
		if len(docs) > 1 {
			path, err = view.RunFileSelector(docs)
			if err != nil {
				return err
			}
		} else if len(docs) < 1 {
			fmt.Printf("The desired file could not be found.\n")
			return nil
		} else {
			path = docs[0]
		}

		doc, err := loader.LoadFile(path)
		if err != nil {
			return err
		}
		view.ViewDocument(doc)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(catCommand)
}
