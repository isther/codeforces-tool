package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get list of contest",
	Long:  "Get the list of questions and the status of the questions",
	Run:   func(cmd *cobra.Command, args []string) {},
}
