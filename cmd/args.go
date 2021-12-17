package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/28251536/codeforces-tool/client"
	"github.com/28251536/codeforces-tool/config"
)

type ParSeArgs struct {
	Info client.Info
}

var Args ParSeArgs

func ParseArgs() error {
	cfg := config.Instance
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	info := &client.Info{}

	if info.ProblemType == "" {
		parsed := parsePath(path)
		if value, ok := parsed["problemType"]; ok {
			info.ProblemType = value
		}
		if value, ok := parsed["contestID"]; ok && info.ContestID == "" {
			info.ContestID = value
		}
		if value, ok := parsed["groupID"]; ok && info.GroupID == "" {
			info.GroupID = value
		}
		if value, ok := parsed["problemID"]; ok && info.ProblemID == "" {
			info.ProblemID = value
		}
	}
	if info.ProblemType == "" || info.ProblemType == "contest" {
		if len(info.ContestID) < 6 {
			info.ProblemType = "contest"
		}
	}

	root := cfg.FolderName["root"]
	info.RootPath = filepath.Join(path, root)
	for {
		base := filepath.Base(path)
		if base == root {
			info.RootPath = path
			break
		}
		if filepath.Dir(path) == path {
			break
		}
		path = filepath.Dir(path)
	}
	info.RootPath = filepath.Join(info.RootPath, cfg.FolderName[info.ProblemType])

	Args.Info = *info

	return nil
}

// ProblemRegStr problem
const ProblemRegStr = `\w+`

// StrictProblemRegStr strict problem
const StrictProblemRegStr = `[a-zA-Z]+\d*`

// ContestRegStr contest
const ContestRegStr = `\d+`

// ArgTypePathRegStr path
var ArgTypePathRegStr = [...]string{
	fmt.Sprintf("%v/%v/((?P<contestID>%v)/((?P<problemID>%v)/)?)?", "%v", "%v", ContestRegStr, ProblemRegStr),
	fmt.Sprintf("%v/%v/((?P<contestID>%v)/((?P<problemID>%v)/)?)?", "%v", "%v", ContestRegStr, ProblemRegStr),
	fmt.Sprintf("%v/%v/((?P<problemID>%v)/)?", "%v", "%v", ProblemRegStr),
}

func parsePath(path string) map[string]string {
	path = filepath.ToSlash(path) + "/"

	output := make(map[string]string)
	cfg := config.Instance

	for k, problemType := range client.ProblemTypes {
		reg := regexp.MustCompile(fmt.Sprintf(ArgTypePathRegStr[k], cfg.FolderName["root"], cfg.FolderName[problemType]))
		names := reg.SubexpNames()
		for i, val := range reg.FindStringSubmatch(path) {
			if names[i] != "" && val != "" {
				output[names[i]] = val
			}
			output["problemType"] = problemType
		}
	}

	return output
}
