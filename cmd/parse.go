package cmd

import (
	"github.com/28251536/codeforces-tool/client"
)

// var parseCmd = &cobra.Command{
// 	Use:   "parse",
// 	Short: "Parse problems",
// 	Long:  "Parse every problems of the contest",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		Parse()
// 	},
// }

func Parse() (err error) {
	cln := client.Instance

	info := Args.Info
	info.ContestID = contestID

	work := func() error {
		_, _, err = cln.Parse(info)
		if err != nil {
			return err
		}

		return nil
	}
	if err = work(); err != nil {
		if err = loginAgain(cln, err); err == nil {
			err = work()
		}
	}
	return
}
