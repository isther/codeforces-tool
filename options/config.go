package options

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var configList = []string{
	"login",
	"add a template",
	"delete a template",
	"set default template",
}

type config struct {
	cursor int
	choice chan string
}

func (m config) Init() tea.Cmd {
	return nil
}

func SelectConfig() string {
	// This is where we'll listen for the choice the user makes in the Bubble
	// Tea program.
	result := make(chan string, 1)
	// defer close(result)

	// Pass the channel to the initialize function so our Bubble Tea program
	// can send the final choice along when the time comes.
	p := tea.NewProgram(config{cursor: 0, choice: result})
	if err := p.Start(); err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Print out the final choice.
	r := <-result
	if r != "" {
		fmt.Printf("\n---\nThe language you choose is %s!\n", r)
	}

	return r
}

func (c config) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			close(c.choice) // If we're quitting just close the channel.
			return c, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			c.choice <- configList[c.cursor]
			return c, tea.Quit

		case "down", "j":
			c.cursor++
			if c.cursor >= len(configList) {
				c.cursor = 0
			}

		case "up", "k":
			c.cursor--
			if c.cursor < 0 {
				c.cursor = len(configList) - 1
			}
		}

	}

	return c, nil
}

func (c config) View() string {
	s := strings.Builder{}
	s.WriteString("What kind of Bubble Tea would you like to order?\n\n")

	for i := 0; i < len(configList); i++ {
		if c.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(configList[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}
