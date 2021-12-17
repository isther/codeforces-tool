package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/28251536/codeforces-tool/client"
	"github.com/28251536/codeforces-tool/config"
	"github.com/28251536/codeforces-tool/options"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	version = "v0.0.1"

	rootCmd = &cobra.Command{
		Use:     "cf",
		Version: version,
		Short:   "A cmd tool for codeforces contest",
		Long: `This is a cmd tool for codeforces contest, 
		it will be help you to create direcory 、download example of problem、
		test for problem、submit problem...
		Good lucky!`,
	}
)

var contestID string
var problemType = "contest"

func init() {
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(raceCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(submitCmd)
	rootCmd.AddCommand(listCmd)

	// testCmd.Flags().IntVarP(&n, "number", "n", 1, "input number")
	raceCmd.Flags().StringVarP(&contestID, "contestId", "c", "", "The id of the match you want to parse")
}

func Execute() error {
	return rootCmd.Execute()
}

func loginAgain(cln *client.Client, err error) error {
	if err != nil && err.Error() == client.ErrorNotLogged {
		color.Red("Not logged. Try to login\n")
		err = cln.Login()
	}
	return err
}

func getSampleID() (samples []string) {
	path, err := os.Getwd()
	if err != nil {
		return
	}

	path = filepath.Join(path, "/Tests")
	paths, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	reg := regexp.MustCompile(`input_(\d+).txt`)
	for _, path := range paths {
		name := path.Name()
		tmp := reg.FindSubmatch([]byte(name))
		if tmp != nil {
			idx := string(tmp[1])
			ans := fmt.Sprintf("./Tests/output_%v.txt", idx)
			if _, err := os.Stat(ans); err == nil {
				samples = append(samples, idx)
			}
		}
	}
	return
}

// CodeList Name matches some template suffix, index are template array indexes
type CodeList struct {
	Name  string
	Index []int
}

func getCode(templates []config.CodeTemplate) (codes []CodeList, err error) {
	mp := make(map[string][]int)
	for i, temp := range templates {
		suffixMap := map[string]bool{}
		for _, suffix := range temp.Suffix {
			if _, ok := suffixMap[suffix]; !ok {
				suffixMap[suffix] = true
				sf := "." + suffix
				mp[sf] = append(mp[sf], i)
			}
		}
	}

	path, err := os.Getwd()
	if err != nil {
		return
	}
	paths, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	for _, path := range paths {
		name := path.Name()
		ext := filepath.Ext(name)
		if idx, ok := mp[ext]; ok {
			codes = append(codes, CodeList{name, idx})
		}
	}

	return codes, nil
}

func getOneCode(templates []config.CodeTemplate) (name string, index int, err error) {
	codes, err := getCode(templates)
	if err != nil {
		return
	}
	if len(codes) < 1 {
		return "", 0, errors.New("Cannot find any code.\nMaybe you should add a new template by `cf config`")
	}
	if len(codes) > 1 {
		color.Cyan("There are multiple files can be selected.")
		for i, code := range codes {
			fmt.Printf("%3v: %v\n", i, code.Name)
		}

		i := -1
		for {
			i = options.ChooseIndex()
			if 0 < i && i < len(codes) {
				break
			}

		}
		codes[0] = codes[i]
	}
	if len(codes[0].Index) > 1 {
		color.Cyan("There are multiple languages match the file.")
		for i, idx := range codes[0].Index {
			fmt.Printf("%3v: %v\n", i, client.Langs[templates[idx].Lang])
		}

		i := -1
		for {
			i = options.ChooseIndex()
			if 0 < i && i < len(codes[0].Index) {
				break
			}

		}

		codes[0].Index[0] = codes[0].Index[i]
	}
	return codes[0].Name, codes[0].Index[0], nil
}
