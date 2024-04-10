package ui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

func OpenTextArea() string {
	p := tea.NewProgram(initialModel())

	m, err := p.Run()

	if err != nil {
		log.Fatal(err)
	}

	return m.(model).textarea.Value()
}

type errMsg error

type model struct {
	textarea textarea.Model
	err      error
}

func initialModel() model {
	ti := textarea.New()
	ti.Placeholder = "Type a query here..."
	ti.SetWidth(80)
	ti.SetHeight(20)
	ti.Focus()

	return model{
		textarea: ti,
		err:      nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	var clearInputs bool = false

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}

		case tea.KeyCtrlC:
			clearInputs = true
			cmds = append(cmds, tea.Quit)

		case tea.KeyCtrlE:
			return m, tea.Quit

		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)

	if clearInputs {
		m.textarea.SetValue("")
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		"\n\n%s\n\n%s",
		m.textarea.View(),
		"(ctrl+e to execute)",
	) + "\n\n"
}
