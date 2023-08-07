package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use: "list",
	Short: "",
	Long: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("docat list")
		return nil
	},
}

func init(){
	rootCmd.AddCommand(listCommand)
}