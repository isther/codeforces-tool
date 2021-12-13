package cmd

import "github.com/spf13/cobra"

var genCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a source file based on the template",
	Long:  "Create a source file based on the template",
	Run:   func(cmd *cobra.Command, args []string) {},
}
