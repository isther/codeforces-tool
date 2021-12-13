package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

var (
	rootCmd = &cobra.Command{
		Use:   "cf",
		Short: "A cmd tool for codeforces contest",
		Long: `This is a cmd tool for codeforces contest, 
		it will be help you to create direcory 、download example of problem、
		test for problem、submit problem...
		Good lucky!`,
	}

	Conf       *Config
	configPath = "config.yaml"
)

func init() {
	setting, err := readConfig()
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}

	err = setting.vp.UnmarshalKey("Config", &Conf)
	if err != nil {
		log.Fatalf("unmarshal config error: %v", err)
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(raceCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(submitCmd)
	rootCmd.AddCommand()

}

func Execute() error {
	return rootCmd.Execute()
}

func readConfig() (*Setting, error) {
	vp := viper.New()
	// vp.SetConfigName("config")
	// vp.AddConfigPath("config/")
	// vp.SetConfigFile("yaml")
	vp.SetConfigFile(configPath)

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
