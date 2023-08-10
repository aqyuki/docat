package commands

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	IgnoreSubDirectoryFlagMsg = "Search only the root directory and do not search sub directories"
)

var (
	IgnoreSubDirectoryFlag bool
)

var rootCmd = &cobra.Command{
	Use:   "docat",
	Short: "Docat provides quick access to documents within a project.",
	Long:  "Docat provides quick access to documents within a project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&IgnoreSubDirectoryFlag, "root", "r", false, IgnoreSubDirectoryFlagMsg)
}
