package options

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type templatePath struct {
	textInput textinput.Model
	eg        string
	err       error

	choice chan string
}

var egList = map[string]string{
	"langs":        `e.g. "42" `,
	"templatepath": `e.g. "~/home/ther/codeforces/template/template.cpp" `,
	"suffix":       `e.g. "cxx cc"`,
	"alias":        `e.g. "cpp" "py"`,
	"email":        `e.g. "28251536@qq.com"`,
	"password":     `e.g. "1234567890"`,
	"beforeScript": `e.g. "g++ $%full%$ -o $%file%$.exe -std=c++11"), empty is ok:`,
	"script":       `"./$%file%$.exe" "python3 $%full%$"):`,
	"afterScript":  `"rm $%file%$.exe" or "cmd.exe /C del $%file%$.exe" in windows), empty is ok:`,
}

var tipList = map[string]string{
	"langs":        `Please enter index of language`,
	"templatepath": `Please enter the path of the template file`,
	"suffix": `The suffix of template above will be added by default.
Other suffix?, empty is ok`,
	"alias":        `Template's alias`,
	"email":        `Please enter your email`,
	"password":     `Please enter your password`,
	"beforeScript": `Please enter before script`,
	"script":       `Please enter script`,
	"afterScript":  `Please enter after script`,
}

func ChooseString(eg string) string {
	ti := textinput.NewModel()
	ti.Placeholder = egList[eg]
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	result := make(chan string, 1)

	p := tea.NewProgram(templatePath{
		textInput: ti,
		choice:    result,
		eg:        eg,
		err:       nil,
	})
	if err := p.Start(); err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
	r := <-result
	// if r != "" {
	// 	fmt.Printf("The path of you template is %s!\n", r)
	// }
	return r
}

func (m templatePath) Init() tea.Cmd {
	return textinput.Blink
}

func (m templatePath) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.choice <- m.textInput.Value()
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m templatePath) View() string {
	tip := tipList[m.eg]
	return fmt.Sprintf(
		"%s: \n\n%s\n\n%s",
		tip,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
