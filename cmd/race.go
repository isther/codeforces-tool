package cmd

import (
	"fmt"
	"time"

	"github.com/28251536/codeforces-tool/client"
	"github.com/28251536/codeforces-tool/config"
	"github.com/spf13/cobra"
)

var raceCmd = &cobra.Command{
	Use:   "race",
	Short: "Parsing questions",
	Long:  "Parsing questions of the contest, download the sample to the local",
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(contestID)
		Race()
	},
}

func Race() (err error) {
	cfg := config.Instance
	cln := client.Instance

	info := Args.Info
	info.ContestID = contestID

	if err = cln.RaceContest(info); err != nil {
		if err = loginAgain(cln, err); err == nil {
			err = cln.RaceContest(info)
			if err != nil {
				return
			}
		}
	}

	time.Sleep(1)
	URL, err := info.ProblemSetURL(cfg.Host)
	if err != nil {
		return
	}
	openURL(URL)
	openURL(URL + "/problems")
	return Parse()
}
