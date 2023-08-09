package commands

import (
	"github.com/aqyuki/docat/internal/scanner"
	"github.com/aqyuki/docat/internal/utils"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Lists the documents contained in the folder.",
	Long:  "Lists the documents contained in the folder. If the folder contains subfolders, the search is extended to the subfolders. The document formats displayed are README, LICENSE, CHANGELOG, CONTRIBUTING, or CONTRIBUTOR.",
	RunE: func(cmd *cobra.Command, args []string) error {

		targetDir, err := utils.TargetDirectoryParser(args)
		if err != nil {
			return err
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
