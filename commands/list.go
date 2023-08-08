package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aqyuki/docat/internal/scanner"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Lists the documents contained in the folder.",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("docat list")

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

		err = scanner.ShowDocumentList(targetDir)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
