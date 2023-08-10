package commands

import (
	"fmt"

	"github.com/aqyuki/docat/internal/display"
	"github.com/aqyuki/docat/internal/document"
	"github.com/aqyuki/docat/internal/loader"
	"github.com/aqyuki/docat/internal/scanner"
	"github.com/aqyuki/docat/internal/scanner/options"
	"github.com/aqyuki/docat/internal/utils"
	"github.com/spf13/cobra"
)

var catCommand = &cobra.Command{
	Use:   "cat",
	Short: "Displays the contents of the selected file.",
	Long:  "Displays the contents of the specified file. The available document formats are README, LICENSE, CHANGELOG, CONTRIBUTING, CONTRIBUTOR or CONTRIBUTE. If there are multiple files of the same document format, a separate selector will be used to open a new file.",
	RunE: func(cmd *cobra.Command, args []string) error {

		opts := options.CreateDefaultOption()
		if IgnoreSubDirectoryFlag {
			options.WithIgnoreSubDirectory(opts)
		}

		targetDir, err := utils.TargetDirectoryParser(args)
		if err != nil {
			return err
		}

		tag, err := display.DocumentSelector()
		if err != nil {
			return err
		}

		files, err := scanner.ScanWithSpinner(targetDir, opts)
		if err != nil {
			return err
		}

		pattern := document.FetchPatternForTag(&tag)
		if pattern == nil {
			return nil
		}

		docs := make([]string, 0)
		for _, file := range files {
			if pattern.Match(file) {
				docs = append(docs, file)
			}
		}

		// docs contains the file paths of the detected documents.
		// If the number of file paths for a document is one, the file is previewed as is,
		// but if there are multiple file paths, a separate file selector is displayed.
		var path string
		if len(docs) > 1 {
			path, err = display.FileSelector(docs)
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
		display.DocumentViewport(doc)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(catCommand)
}
