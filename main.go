package main

import (
	"os"

	"github.com/28251536/codeforces-tool/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
