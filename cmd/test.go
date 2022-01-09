package cmd

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unicode"

	"github.com/28251536/codeforces-tool/client"
	"github.com/28251536/codeforces-tool/config"
	"github.com/28251536/codeforces-tool/util"
	"github.com/fatih/color"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/shirou/gopsutil/process"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test samples",
	Long:  "To test the samples given by the question",
	Run: func(cmd *cobra.Command, args []string) {
		err := Test()
		if err != nil {
			color.Red("Test failed")
		}
	},
}

// Test command
func Test() (err error) {
	cfg := config.Instance
	if len(cfg.Template) == 0 {
		return errors.New("You have to add at least one code template by `cf config`")
	}
	samples := getSampleID()
	if len(samples) == 0 {
		return errors.New("Cannot find any sample file")
	}
	filename, index, err := getOneCode(cfg.Template)
	if err != nil {
		return
	}
	template := cfg.Template[index]
	path, full := filepath.Split(filename)
	ext := filepath.Ext(filename)
	file := full[:len(full)-len(ext)]
	rand := util.RandString(8)

	filter := func(cmd string) string {
		cmd = strings.ReplaceAll(cmd, "$%rand%$", rand)
		cmd = strings.ReplaceAll(cmd, "$%path%$", path)
		cmd = strings.ReplaceAll(cmd, "$%full%$", full)
		cmd = strings.ReplaceAll(cmd, "$%file%$", file)
		return cmd
	}

	run := func(script string) error {
		if s := filter(script); len(s) > 0 {
			cmds := splitCmd(s)
			cmd := exec.Command(cmds[0], cmds[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		}
		return nil
	}

	if err = run(template.BeforeScript); err != nil {
		return
	}

	var limit client.Limit
	util.Load("./.config",&limit)

	if s := filter(template.Script); len(s) > 0 {
		for _, i := range samples {
			err := judge(i, s, time.Duration(limit.TimeLimit*1000*uint64(time.Millisecond)))
			if err != nil {
				color.Red(err.Error())
			}
		}
	} else {
		return errors.New("Invalid script command. Please check config file")
	}
	return run(template.AfterScript)
}

func splitCmd(s string) (res []string) {
	// https://github.com/vrischmann/shlex/blob/master/shlex.go
	var buf bytes.Buffer
	insideQuotes := false
	for _, r := range s {
		switch {
		case unicode.IsSpace(r) && !insideQuotes:
			if buf.Len() > 0 {
				res = append(res, buf.String())
				buf.Reset()
			}
		case r == '"' || r == '\'':
			if insideQuotes {
				res = append(res, buf.String())
				buf.Reset()
				insideQuotes = false
				continue
			}
			insideQuotes = true
		default:
			buf.WriteRune(r)
		}
	}
	if buf.Len() > 0 {
		res = append(res, buf.String())
	}
	return
}

func plain(raw []byte) string {
	buf := bufio.NewScanner(bytes.NewReader(raw))
	var b bytes.Buffer
	newline := []byte{'\n'}
	for buf.Scan() {
		b.Write(bytes.TrimSpace(buf.Bytes()))
		b.Write(newline)
	}
	return b.String()
}

func judge(sampleID, command string, duration time.Duration) error {
	inPath := fmt.Sprintf("./Tests/input_%v.txt", sampleID)
	ansPath := fmt.Sprintf("./Tests/output_%v.txt", sampleID)
	input, err := os.Open(inPath)
	if err != nil {
		return err
	}
	var o bytes.Buffer
	output := io.Writer(&o)

	cmds := splitCmd(command)

	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Stdin = input
	cmd.Stdout = output
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Runtime Error #%v ... %v", sampleID, err.Error())
	}

	pid := cmd.Process.Pid
	maxMemory := uint64(0)
	ch := make(chan error)
	go func() {
		ch <- cmd.Wait()
	}()

	ctx := context.Background()
	ctxWithTimeOut, ctxCancleFunc := context.WithTimeout(ctx, duration)
	defer ctxCancleFunc()

	running := true
	for running {
		select {
		case err := <-ch:
			if err != nil {
				return fmt.Errorf("Runtime Error #%v ... %v", sampleID, err.Error())
			}
			running = false
		case <-ctxWithTimeOut.Done():
			_ = syscall.Kill(pid, syscall.SIGKILL)
			return fmt.Errorf("Time out")
		default:
			p, err := process.NewProcess(int32(pid))
			if err == nil {
				m, err := p.MemoryInfo()
				if err == nil && m.RSS > maxMemory {
					maxMemory = m.RSS
				}
			}
		}
	}

	b, err := ioutil.ReadFile(ansPath)
	if err != nil {
		b = []byte{}
	}
	ans := plain(b)
	out := plain(o.Bytes())

	state := ""
	diff := ""
	if out == ans {
		state = color.New(color.FgGreen).Sprintf("Passed #%v", sampleID)
	} else {
		input, err := ioutil.ReadFile(inPath)
		if err != nil {
			return err
		}
		state = color.New(color.FgRed).Sprintf("Failed #%v", sampleID)
		dmp := diffmatchpatch.New()
		d := dmp.DiffMain(out, ans, true)
		diff += color.New(color.FgCyan).Sprintf("-----Input-----\n")
		diff += string(input) + "\n"
		diff += color.New(color.FgCyan).Sprintf("-----Output-----\n")
		diff += dmp.DiffText1(d) + "\n"
		diff += color.New(color.FgCyan).Sprintf("-----Answer-----\n")
		diff += dmp.DiffText2(d) + "\n"
		diff += color.New(color.FgCyan).Sprintf("-----Diff-----\n")
		diff += dmp.DiffPrettyText(d) + "\n"
	}

	parseMemory := func(memory uint64) string {
		if memory > 1024*1024 {
			return fmt.Sprintf("%.3fMB", float64(memory)/1024.0/1024.0)
		} else if memory > 1024 {
			return fmt.Sprintf("%.3fKB", float64(memory)/1024.0)
		}
		return fmt.Sprintf("%vB", memory)
	}

	fmt.Printf("%v ... %.3fs %v\n%v", state, cmd.ProcessState.UserTime().Seconds(), parseMemory(maxMemory), diff)
	return nil
}
