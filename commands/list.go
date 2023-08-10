package commands

import (
	"github.com/aqyuki/docat/internal/scanner"
	scop "github.com/aqyuki/docat/internal/scanner/options"
	"github.com/aqyuki/docat/internal/utils"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "Lists the documents contained in the folder.",
	Long:  "Lists the documents contained in the folder. If the folder contains sub directory, the search is extended to the subfolders. The document formats displayed are README, LICENSE, CHANGELOG, CONTRIBUTING, CONTRIBUTOR or CONTRIBUTE.",
	RunE: func(cmd *cobra.Command, args []string) error {

		scannerOpts := scop.CreateDefaultOption()
		if IgnoreSubDirectoryFlag {
			scop.WithIgnoreSubDirectory(scannerOpts)
		}

		targetDir, err := utils.TargetDirectoryParser(args)
		if err != nil {
			return err
		}

		err = scanner.ShowDocumentList(targetDir, scannerOpts)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCommand)
}
