package commands

import (
	"fmt"
	"os"

	"github.com/aqyuki/docat/internal/scanner"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "",
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
			targetDir = args[0]
		}

		err = scanner.ScanDocument(targetDir)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
