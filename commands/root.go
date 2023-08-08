package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docat",
	Short: "Docat provides quick access to documents within a project.",
	Long:  "Docat provides quick access to documents within a project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("main command")
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
}
