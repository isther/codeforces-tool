package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/28251536/codeforces-tool/client"
	"github.com/28251536/codeforces-tool/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Create a source file based on the template",
	Long:  "Create a source file based on the template",
	Run: func(cmd *cobra.Command, args []string) {
		err := Gen()
		if err != nil {
			color.Red("Generate failed")
		}
	},
}

func Gen() (err error) {
	cfg := config.Instance
	if len(cfg.Template) == 0 {
		return errors.New("You have to add at least one code template by `cf config`")
	}

	path := cfg.Template[cfg.Default].Path
	cln := client.Instance

	fmt.Println(1)
	source, err := readTemplateSource(path, cln)
	if err != nil {
		color.Red("Read template source filed")
		return
	}

	currentPath, err := os.Getwd()
	if err != nil {
		return
	}

	ext := filepath.Ext(path)
	return gen(source, currentPath, ext)
}

func readTemplateSource(path string, cln *client.Client) (source string, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	source = parseTemplate(string(b), cln)
	return
}

func parseTemplate(source string, cln *client.Client) string {
	now := time.Now()
	source = strings.ReplaceAll(source, "$%U%$", cln.Handle)
	source = strings.ReplaceAll(source, "$%Y%$", fmt.Sprintf("%v", now.Year()))
	source = strings.ReplaceAll(source, "$%M%$", fmt.Sprintf("%02v", int(now.Month())))
	source = strings.ReplaceAll(source, "$%D%$", fmt.Sprintf("%02v", now.Day()))
	source = strings.ReplaceAll(source, "$%h%$", fmt.Sprintf("%02v", now.Hour()))
	source = strings.ReplaceAll(source, "$%m%$", fmt.Sprintf("%02v", now.Minute()))
	source = strings.ReplaceAll(source, "$%s%$", fmt.Sprintf("%02v", now.Second()))
	return source
}

func gen(source, currentPath, ext string) error {
	path := filepath.Join(currentPath, filepath.Base(currentPath))

	savePath := path + ext
	i := 1
	for _, err := os.Stat(savePath); err == nil; _, err = os.Stat(savePath) {
		tmpPath := fmt.Sprintf("%v%v%v", path, i, ext)
		fmt.Printf("%v exists. Rename to %v\n", filepath.Base(savePath), filepath.Base(tmpPath))
		savePath = tmpPath
		i++
	}

	err := ioutil.WriteFile(savePath, []byte(source), 0644)
	if err == nil {
		color.Green("Generated! See %v", filepath.Base(savePath))
	}
	return err
}
