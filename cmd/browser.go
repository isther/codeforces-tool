package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

var commands = map[string]string{
	// "windows": "cmd /c start",
	"windows": "cmd",
	"darwin":  "open",
	"linux":   "xdg-open",
}

// Open calls the OS default program for uri
func openURL(url string) error {
	color.Green("Open %v", url)
	if runtime.GOOS == "windows" {
		url = fmt.Sprintf("/c start %v", url)
	}

	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, url)
	return cmd.Start()
}
