package main

import (
	"github.com/28251536/codeforces-tool/client"
	"github.com/28251536/codeforces-tool/cmd"
	"github.com/28251536/codeforces-tool/config"
)

const configPath = "~/.cfs/config"
const clnPath = "~/.cfs/session"

func init() {
	config.Init(configPath)
	client.Init(clnPath, config.Instance.Host)

	cmd.ParseArgs()
}

func main() {
	cmd.Execute()
}
