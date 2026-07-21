package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProjectListView struct {
	projects []string
	cursor   int
	viewport viewport.Model
	width    int
	height   int
	ready    bool
}

func NewProjectListView() *ProjectListView {
	return &ProjectListView{
		projects: []string{},
		cursor:   0,
	}
}

func (p *ProjectListView) Init() tea.Cmd {
	return nil
}

func (p *ProjectListView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if p.cursor < len(p.projects)-1 {
				p.cursor++
			}
		case "k", "up":
			if p.cursor > 0 {
				p.cursor--
			}
		case "enter":
			if len(p.projects) > 0 {
				return p, nil
			}
		case "esc":
			return p, nil
		}

	case tea.WindowSizeMsg:
		p.width = msg.Width
		p.height = msg.Height
		p.viewport.Width = msg.Width
		p.viewport.Height = msg.Height - 4
		p.ready = true
	}

	if p.ready {
		p.viewport.SetContent(p.renderContent())
	}

	return p, nil
}

func (p *ProjectListView) View() string {
	if !p.ready {
		return "Lade Projekte..."
	}

	title := lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1).
		Render("Projekt laden")

	content := p.viewport.View()

	return lipgloss.JoinVertical(lipgloss.Top, title, content)
}

func (p *ProjectListView) renderContent() string {
	var sb strings.Builder

	for i, project := range p.projects {
		cursor := "  "
		if i == p.cursor {
			cursor = "> "
		}
		sb.WriteString(fmt.Sprintf("%s%s\n", cursor, project))
	}

	if len(p.projects) == 0 {
		sb.WriteString("Keine Projekte gefunden")
	}

	return sb.String()
}

func (p *ProjectListView) GetShortcuts() string {
	return "[↑]↑ [↓]↓ [Enter]Auswählen [Esc]Zurück"
}

func (p *ProjectListView) SetProjects(projects []string) {
	p.projects = projects
}

func (p *ProjectListView) GetSelectedProject() string {
	if len(p.projects) > 0 && p.cursor >= 0 && p.cursor < len(p.projects) {
		return p.projects[p.cursor]
	}
	return ""
}
