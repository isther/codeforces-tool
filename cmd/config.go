package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Config struct {
	Name string
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config file",
	Long:  "Config file of tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello, this is config")
		fmt.Printf("name = %v", Conf.Name)
	},
}
