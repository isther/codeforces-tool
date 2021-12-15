package cmd

import (
	"github.com/28251536/codeforces-tool/client"
	"github.com/28251536/codeforces-tool/config"
	"github.com/28251536/codeforces-tool/options"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config file",
	Long:  "Config file of tool",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Instance
		cln := client.Instance
		op := options.SelectConfig()
		switch op {
		case "login":
			cln.ConfigLogin()
		case "add a template":
			cfg.AddTemplate()
		case "delete a template":
			cfg.RemoveTemplate()
		case "set default template":
			cfg.SetDefaultTemplate()
		}
	},
}
