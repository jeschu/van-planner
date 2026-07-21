package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DeleteDialogView struct {
	width        int
	height       int
	confirmIndex int
}

func NewDeleteDialogView() *DeleteDialogView {
	return &DeleteDialogView{
		confirmIndex: 0,
	}
}

func (d *DeleteDialogView) Init() tea.Cmd {
	return nil
}

func (d *DeleteDialogView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "right", "tab":
			d.confirmIndex = 1 - d.confirmIndex
		case "enter", " ":
			if d.confirmIndex == 0 {
				return d, func() tea.Msg {
					return ConfirmDeleteMsg{}
				}
			}
			return d, func() tea.Msg {
				return CancelDeleteMsg{}
			}
		case "esc":
			return d, func() tea.Msg {
				return CancelDeleteMsg{}
			}
		}
	}

	return d, nil
}

func (d *DeleteDialogView) View() string {
	question := "Produkt wirklich löschen?"

	var yesBtn, noBtn string
	if d.confirmIndex == 0 {
		yesBtn = dialogButtonStyle.Render("Ja")
		noBtn = dialogButtonInactiveStyle.Render("Nein")
	} else {
		yesBtn = dialogButtonInactiveStyle.Render("Ja")
		noBtn = dialogButtonStyle.Render("Nein")
	}

	buttons := lipgloss.JoinHorizontal(lipgloss.Center, yesBtn, "  ", noBtn)
	content := lipgloss.JoinVertical(lipgloss.Center, question, "\n", buttons)

	return dialogBoxStyle.Render(content)
}

func (d *DeleteDialogView) GetShortcuts() string {
	return "[←][→] Auswahl [Enter] Bestätigen [Esc] Abbrechen"
}

type ConfirmDeleteMsg struct{}

type CancelDeleteMsg struct{}
