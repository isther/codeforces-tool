package cmd

import "github.com/spf13/cobra"

var raceCmd = &cobra.Command{
	Use:   "race",
	Short: "Parsing questions",
	Long:  "Parsing questions of the contest, download the sample to the local",
	Run:   func(cmd *cobra.Command, args []string) {},
}
