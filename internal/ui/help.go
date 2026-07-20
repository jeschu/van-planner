package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type HelpKey struct {
	Key  string
	Help string
}

func GetHelpKeyMap() []HelpKey {
	return []HelpKey{
		{"j/↓", "Nächstes Produkt"},
		{"k/↑", "Vorheriges Produkt"},
		{"Space", "Completed toggle"},
		{"n", "Neues Produkt"},
		{"e", "Produkt editieren"},
		{"d", "Produkt löschen"},
		{"/", "Suche"},
		{"K", "Neue Kategorie"},
		{"E", "Kategorie editieren"},
		{"D", "Kategorie löschen"},
		{"Ctrl+O", "Projekt öffnen"},
		{"Ctrl+S", "Projekt speichern"},
		{"?", "Hilfe"},
		{"q", "Beenden"},
		{"Ctrl+C", "Beenden"},
	}
}

func GetContextualHelp(mode mode) string {
	switch mode {
	case modeList:
		return "j/k: Navigation | Space: Toggle | n: Neu | e: Edit | d: Delete | /: Suche | ?: Hilfe | q: Quit"
	case modeCreate, modeEdit:
		return "Tab/Enter: Weiter | Shift+Tab: Zurück | Enter: Speichern | Esc: Abbrechen"
	case modeCreateCategory, modeEditCategory:
		return "Enter: Speichern | Esc: Abbrechen"
	case modeProjectList:
		return "Enter: Laden | n: Neu | d: Löschen | q: Quit"
	default:
		return ""
	}
}

func RenderHelpPage() string {
	var b strings.Builder

	b.WriteString(TitleStyle.Render("Hilfe – Alle Tastatur-Shortcuts") + "\n\n")

	b.WriteString(renderKeySection("Navigation", []HelpKey{
		{"j/↓", "Nächstes Produkt"},
		{"k/↑", "Vorheriges Produkt"},
		{"Space", "Completed toggle"},
	}))

	b.WriteString(renderKeySection("Aktionen", []HelpKey{
		{"n", "Neues Produkt"},
		{"e", "Produkt editieren"},
		{"d", "Produkt löschen"},
		{"/", "Suche"},
	}))

	b.WriteString(renderKeySection("Kategorien", []HelpKey{
		{"K", "Neue Kategorie"},
		{"E", "Kategorie editieren"},
		{"D", "Kategorie löschen"},
	}))

	b.WriteString(renderKeySection("Projekte", []HelpKey{
		{"Ctrl+O", "Projekt öffnen"},
		{"Ctrl+S", "Projekt speichern"},
	}))

	b.WriteString(renderKeySection("Allgemein", []HelpKey{
		{"?", "Hilfe"},
		{"q", "Beenden"},
		{"Ctrl+C", "Beenden"},
	}))

	b.WriteString("\n")
	b.WriteString(HelpStyle.Render("Drücke eine beliebige Taste zum Schließen"))
	b.WriteString("\n")

	return b.String()
}

func renderKeySection(title string, keys []HelpKey) string {
	var b strings.Builder

	b.WriteString(SubtitleStyle.Render(title) + "\n")

	maxKeyLen := 0
	for _, k := range keys {
		if len(k.Key) > maxKeyLen {
			maxKeyLen = len(k.Key)
		}
	}

	for _, k := range keys {
		b.WriteString(fmt.Sprintf("  %-*s  %s\n", maxKeyLen+2, k.Key+":", k.Help))
	}

	b.WriteString("\n")

	return b.String()
}

type helpModel struct {
	quitting bool
}

func newHelpModel() helpModel {
	return helpModel{
		quitting: false,
	}
}

func (m helpModel) Init() tea.Cmd {
	return nil
}

func (m helpModel) Update(msg tea.Msg) (helpModel, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		m.quitting = true
		return m, tea.Quit
	}
	return m, nil
}

func (m helpModel) View() string {
	return RenderHelpPage()
}

func (m helpModel) IsQuitting() bool {
	return m.quitting
}
