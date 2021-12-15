package options

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type index struct {
	textInput textinput.Model
	err       error

	choice chan string
}

func ChooseIndex() int {
	ti := textinput.NewModel()
	ti.Placeholder = fmt.Sprintf(`e.g. "1" `)
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	result := make(chan string, 1)

	p := tea.NewProgram(index{
		textInput: ti,
		choice:    result,
		err:       nil,
	})
	if err := p.Start(); err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
	r := <-result
	// if r != "" {
	// 	fmt.Printf("The index of the template is %s!\n", r)
	// }

	idx, _ := strconv.ParseInt(r, 0, 0)
	return int(idx)
}

func (m index) Init() tea.Cmd {
	return textinput.Blink
}

func (m index) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m index) View() string {
	return fmt.Sprintf(
		"Template absolute path: \n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
