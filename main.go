package main

import (
	"fmt"

	"github.com/28251536/codeforces-tool/client"
	"github.com/28251536/codeforces-tool/cmd"
	"github.com/28251536/codeforces-tool/config"
	"github.com/28251536/codeforces-tool/util"
)

const configPath = "/.cf/config"
const clientPath = "/.cf/session"

func init() {
	home, _ := util.Home()
	cfgPath := fmt.Sprintf("%v%v", home, configPath)
	clnPath := fmt.Sprintf("%v%v", home, clientPath)
	// fmt.Printf("cfg = %v", cfgPath)
	// fmt.Printf("cln = %v", clnPath)

	config.Init(cfgPath)
	client.Init(clnPath, config.Instance.Host)

	cmd.ParseArgs()
}

func main() {
	cmd.Execute()
}
