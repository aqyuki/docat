package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	Version = "1.2.0"
)

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Show docat version",
	Long:  "Show docat version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
