package cmd

import "github.com/spf13/cobra"

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test sample",
	Long:  "To test the sample given by the question",
	Run:   func(cmd *cobra.Command, args []string) {},
}
