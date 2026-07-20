package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type categoryFormModel struct {
	input      textinput.Model
	submitting bool
	mode       string
}

func newCategoryFormModel(mode string, editCategory string) categoryFormModel {
	input := textinput.New()
	input.Placeholder = "Kategoriename"
	input.Focus()
	input.CharLimit = 50
	input.Width = 40

	if editCategory != "" {
		input.SetValue(editCategory)
	}

	return categoryFormModel{
		input:      input,
		submitting: false,
		mode:       mode,
	}
}

func (m categoryFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m categoryFormModel) Update(msg tea.Msg) (categoryFormModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.submitting = true
			return m, tea.Quit
		case "enter":
			m.submitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m categoryFormModel) View() string {
	var b strings.Builder

	b.WriteString("Name der Kategorie:\n")
	b.WriteString(m.input.View() + "\n\n")
	b.WriteString(HelpStyle.Render("Enter: Speichern | Esc: Abbrechen"))
	b.WriteString("\n")

	return b.String()
}

func (m categoryFormModel) GetCategoryName() string {
	return strings.TrimSpace(m.input.Value())
}
