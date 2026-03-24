package main

import (
    "log"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	index int
	questions []string
	width int
	height int
	answerField textinput.Model
}

func New(questions []string) *model {
	answerField := textinput.New()

	return &model{questions: questions, answerField: answerField}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.width = msg.Width
			m.height = msg.Height
		case tea.KeyMsg:
			switch msg.String() {
				case "ctrl+c":
					return m, tea.Quit
			}
	}
	return m, nil
}

func (m model) View() string {
	if m.width == 0 {
		return "loading.."
	}
	return lipgloss.JoinVertical(lipgloss.Center, m.questions[m.index], m.answerField.View())
}

func main() {
	questions := []string{"What is your name?", "What is my favorite TUI?", "What is your favorite quote?"}
	m := New(questions)

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer f.Close()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}