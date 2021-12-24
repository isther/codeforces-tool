package cmd

import (
	"os"

	"github.com/28251536/codeforces-tool/client"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var skdCmd = &cobra.Command{
	Use:   "skd",
	Short: "Get the schedule of the contest",
	Long:  "Get the schedule of the contest",
	Run: func(cmd *cobra.Command, args []string) {
		err := Schedule()
		if err != nil {
			color.Red("Get schedule of contest failed")
		}
	},
}

func Schedule() error {
	list := client.IDList{}

	text := list.GetContestList()
	data := [][]string{}
	for _, v := range text {
		data = append(data, []string{
			v.ID,
			v.Name,
			v.Start,
			v.Length,
			v.BeforeStart,
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Start", "Length", "BeforeStart"})

	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetCaption(true, "Good Lucky!")

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
	return nil
}
