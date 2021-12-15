package cmd

import (
	"github.com/spf13/cobra"
)

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit question",
	Long:  "Submit answers to codeforces website",
	Run:   func(cmd *cobra.Command, args []string) {},
}
