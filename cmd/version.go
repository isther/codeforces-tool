package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "v0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of codeforces-tool",
	Long:  `All software has versions. This is codeforces-tool's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
