package cmd

import (
	"io/ioutil"

	"github.com/28251536/codeforces-tool/client"
	"github.com/28251536/codeforces-tool/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit question",
	Long:  "Submit answers to codeforces website",
	Run: func(cmd *cobra.Command, args []string) {
		err := Submit()
		if err != nil {
			color.Red("Submit problem faeild")
		}
	},
}

func Submit() (err error) {
	cln := client.Instance
	cfg := config.Instance
	info := Args.Info

	filename, index, err := getOneCode(cfg.Template)
	if err != nil {
		return
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	source := string(bytes)

	lang := cfg.Template[index].Lang
	if err = cln.Submit(info, lang, source); err != nil {
		if err = loginAgain(cln, err); err == nil {
			err = cln.Submit(info, lang, source)
		}
	}
	return
}
