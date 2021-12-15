package main

import (
	"github.com/28251536/codeforces-tool/client"
	"github.com/28251536/codeforces-tool/cmd"
	"github.com/28251536/codeforces-tool/config"
)

const configPath = "./config.json"
const clnPath = "./session.json"

func init() {
	config.Init(configPath)
	client.Init(clnPath, config.Instance.Host)
}

func main() {
	cmd.Execute()
}
