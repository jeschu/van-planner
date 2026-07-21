package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
Van Planner - Hilfe

Navigation:
  j / ↓         Nächstes Produkt
  k / ↑         Vorheriges Produkt
  ← / →         Kategorie wechseln

Aktionen:
  n             Neues Produkt erstellen
  e             Produkt bearbeiten
  d             Produkt löschen
  Space         Produkt als erledigt markieren
  /             Suche öffnen
  Ctrl-O        Projekt laden
  ?             Hilfe anzeigen

Formular:
  Tab / Enter   Nächstes Feld
  Shift-Tab     Vorheriges Feld
  Esc           Abbrechen
`

	style := lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.RoundedBorder())

	return style.Render(helpText)
}

func (h *HelpView) GetShortcuts() string {
	return "[Esc/?] Hilfe schließen"
}
