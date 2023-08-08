package commands

import "github.com/spf13/cobra"

var catCommand = &cobra.Command{
	Use:   "cat",
	Short: "",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(catCommand)
}
