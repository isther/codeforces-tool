package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/28251536/codeforces-tool/util"

	"github.com/fatih/color"
)

type Limit struct {
	TimeLimit   uint64
	MemoryLimit uint64
}

func findSample(body []byte) (input [][]byte, output [][]byte, err error) {
	irg := regexp.MustCompile(`class="input"[\s\S]*?<pre>([\s\S]*?)</pre>`)
	org := regexp.MustCompile(`class="output"[\s\S]*?<pre>([\s\S]*?)</pre>`)
	a := irg.FindAllSubmatch(body, -1)
	b := org.FindAllSubmatch(body, -1)
	if a == nil || b == nil || len(a) != len(b) {
		return nil, nil, fmt.Errorf("Cannot parse sample with input %v and output %v", len(a), len(b))
	}
	newline := regexp.MustCompile(`<[\s/br]+?>`)
	filter := func(src []byte) []byte {
		src = newline.ReplaceAll(src, []byte("\n"))
		s := html.UnescapeString(string(src))
		return []byte(strings.TrimSpace(s) + "\n")
	}
	for i := 0; i < len(a); i++ {
		input = append(input, filter(a[i][1]))
		output = append(output, filter(b[i][1]))
	}
	return
}

func findProblemLimit(body []byte) (limit *Limit, err error) {
	timeRg := regexp.MustCompile(`time limit per test</div>[\s\S]*?second`)
	memoryRg := regexp.MustCompile(`memory limit per test</div>[\s\S]*?megabytes`)
	a := timeRg.FindAllSubmatch(body, -1)
	b := memoryRg.FindAllSubmatch(body, -1)
	if a == nil || b == nil {
		return nil, fmt.Errorf("Cannot parse timelimit and memorylimit")
	}

	timeLimit := a[0][0]
	memoryLimit := b[0][0]
	timeLimit = bytes.TrimPrefix(timeLimit, []byte("time limit per test</div>"))
	timeLimit = bytes.TrimSuffix(timeLimit, []byte(" second"))
	memoryLimit = bytes.TrimPrefix(memoryLimit, []byte("memory limit per test</div>"))
	memoryLimit = bytes.TrimSuffix(memoryLimit, []byte(" megabytes"))
	time, _ := strconv.ParseUint(string(timeLimit), 10, 64)
	memory, _ := strconv.ParseUint(string(memoryLimit), 10, 64)

	return &Limit{
		TimeLimit:   time,
		MemoryLimit: memory,
	}, nil
}

// ParseProblem parse problem to path. mu can be nil
func (c *Client) ParseProblem(URL, path string, mu *sync.Mutex) (samples int, standardIO bool, err error) {
	body, err := util.GetBody(c.client, URL)
	if err != nil {
		return
	}

	_, err = findHandle(body)
	if err != nil {
		return
	}

	input, output, err := findSample(body)
	if err != nil {
		return
	}

	standardIO = true
	if !bytes.Contains(body, []byte(`<div class="input-file"><div class="property-title">input</div>standard input</div><div class="output-file"><div class="property-title">output</div>standard output</div>`)) {
		standardIO = false
	}

	limit, _ := findProblemLimit(body)
	data, _ := json.MarshalIndent(limit, "", "  ")
	_ = ioutil.WriteFile(filepath.Join(path, "/.config"), data, 0644)

	path = filepath.Join(path, "/Tests")
	if err = os.Mkdir(path, os.ModePerm); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(input); i++ {
		fileIn := filepath.Join(path, fmt.Sprintf("input_%v.txt", i+1))
		fileOut := filepath.Join(path, fmt.Sprintf("output_%v.txt", i+1))

		e := ioutil.WriteFile(fileIn, input[i], 0644)
		if e != nil {
			if mu != nil {
				mu.Lock()
			}
			color.Red(e.Error())
			if mu != nil {
				mu.Unlock()
			}
		}
		e = ioutil.WriteFile(fileOut, output[i], 0644)
		if e != nil {
			if mu != nil {
				mu.Lock()
			}
			color.Red(e.Error())
			if mu != nil {
				mu.Unlock()
			}
		}
	}
	return len(input), standardIO, nil
}

// Parse parse
func (c *Client) Parse(info Info) (problems []string, paths []string, err error) {
	color.Cyan("Parse " + info.Hint())

	problemID := info.ProblemID
	info.ProblemID = "%v"
	urlFormatter, err := info.ProblemURL(c.host)
	if err != nil {
		return
	}
	info.ProblemID = ""
	if problemID == "" {
		statics, err := c.Statis(info)
		if err != nil {
			return nil, nil, err
		}
		problems = make([]string, len(statics))
		for i, problem := range statics {
			problems[i] = problem.ID
		}
	} else {
		problems = []string{problemID}
	}
	contestPath := info.Path()
	fmt.Printf(color.CyanString("The problem(s) will be saved to %v\n"), color.GreenString(contestPath))

	wg := sync.WaitGroup{}
	wg.Add(len(problems))
	mu := sync.Mutex{}
	paths = make([]string, len(problems))
	for i, problemID := range problems {
		paths[i] = filepath.Join(contestPath, strings.ToLower(problemID))
		go func(problemID, path string) {
			defer wg.Done()
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return
			}
			URL := fmt.Sprintf(urlFormatter, problemID)

			samples, standardIO, err := c.ParseProblem(URL, path, &mu)
			if err != nil {
				return
			}

			warns := ""
			if !standardIO {
				warns = color.YellowString("Non standard input output format.")
			}
			mu.Lock()
			if err != nil {
				color.Red("Failed %v. Error: %v", problemID, err.Error())
			} else {
				fmt.Printf("%v %v\n", color.GreenString("Parsed %v with %v samples.", problemID, samples), warns)
			}
			mu.Unlock()
		}(problemID, paths[i])
	}
	wg.Wait()
	return
}
