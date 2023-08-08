package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aqyuki/docat/internal/scanner"
	"github.com/aqyuki/docat/internal/view"
	"github.com/spf13/cobra"
)

var catCommand = &cobra.Command{
	Use:   "cat",
	Short: "",
	Long:  "",
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

		files,err := scanner.CreateDocumentList(targetDir)
		if err != nil {
			return err
		}
		view.RunDocumentSelector(files)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(catCommand)
}
