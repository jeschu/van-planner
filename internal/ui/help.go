package ui

import (
	"github.com/charmbracelet/bubbletea"
)

type HelpView struct {
	width  int
	height int
}

func NewHelpView() *HelpView {
	return &HelpView{}
}

func (h *HelpView) Init() tea.Cmd {
	return nil
}

func (h *HelpView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "?":
			return h, nil
		}

	case tea.WindowSizeMsg:
		h.width = msg.Width
		h.height = msg.Height
	}

	return h, nil
}

func (h *HelpView) View() string {
	helpText := `
☀️ Van Planner - Hilfe

Navigation:
  j / ↓         Nächstes Produkt
  k / ↑         Vorheriges Produkt
  ← / →         Scrollen

Aktionen:
  Space         Produkt als erledigt markieren
  Ctrl-O        Projekt laden
  ?             Hilfe anzeigen

Allgemein:
  q             Beenden
  Esc           Zurück
`

	return helpBorderStyle.Render(helpText)
}

func (h *HelpView) GetShortcuts() string {
	return "[Esc/?] Hilfe schließen"
}
