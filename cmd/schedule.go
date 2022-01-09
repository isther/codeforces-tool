package cmd

import (
	"os"
	"strconv"
	"time"

	"github.com/28251536/codeforces-tool/client"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	skdCmd = &cobra.Command{
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
	format = "2006-01-02 15:04:05"
)

func Schedule() error {
	text := client.GetContest()
	data := [][]string{}
	for _, v := range text {
		data = append(data, []string{
			strconv.FormatUint(v.ID, 10),
			string(v.Name),
			time.Unix(v.StartTimeSeconds, 0).Format(format),
			strconv.FormatInt(v.DurationSeconds/60, 10) + "min",
		})

	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Start", "Length"})

	table.SetRowLine(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetCaption(true, "Good Lucky!")

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
	return nil
}
