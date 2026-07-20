package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jens/van-planner/internal/storage"
)

type projectItem struct {
	info storage.ProjectInfo
}

func (i projectItem) Title() string {
	return i.info.Name
}

func (i projectItem) Description() string {
	return i.info.FilePath
}

func (i projectItem) FilterValue() string {
	return i.info.Name
}

type ProjectListMode int

const (
	ProjectListModeSelect ProjectListMode = iota
	ProjectListModeCreate
	ProjectListModeDelete
)

type projectListModel struct {
	list           list.Model
	projectManager *storage.ProjectManager
	projects       []storage.ProjectInfo
	mode           ProjectListMode
	nameInput      textinput.Model
	selected       *storage.ProjectInfo
	quitting       bool
	message        string
}

func newProjectListModel(pm *storage.ProjectManager) projectListModel {
	projects, _ := pm.ListProjects()

	items := make([]list.Item, len(projects))
	for i, p := range projects {
		items[i] = projectItem{info: p}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Projekt auswählen"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)

	nameInput := textinput.New()
	nameInput.Placeholder = "Projektname"
	nameInput.Focus()
	nameInput.CharLimit = 50
	nameInput.Width = 40

	return projectListModel{
		list:           l,
		projectManager: pm,
		projects:       projects,
		mode:           ProjectListModeSelect,
		nameInput:      nameInput,
	}
}

func (m projectListModel) Init() tea.Cmd {
	return nil
}

func (m projectListModel) Update(msg tea.Msg) (projectListModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch m.mode {
	case ProjectListModeCreate:
		var cmd tea.Cmd
		m.nameInput, cmd = m.nameInput.Update(msg)
		cmds = append(cmds, cmd)

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				m.mode = ProjectListModeSelect
				m.message = "Erstellen abgebrochen"
				return m, tea.Batch(cmds...)
			case "enter":
				name := strings.TrimSpace(m.nameInput.Value())
				if name == "" {
					m.message = "Projektname ist erforderlich"
					return m, tea.Batch(cmds...)
				}
				if m.projectManager.ProjectExists(name) {
					m.message = "Projekt existiert bereits"
					return m, tea.Batch(cmds...)
				}
				if err := m.projectManager.CreateProject(name); err != nil {
					m.message = "Fehler: " + err.Error()
					return m, tea.Batch(cmds...)
				}
				m.selected = &storage.ProjectInfo{
					Name:     name,
					FilePath: fmt.Sprintf("%s/%s.json", storage.ProjectsDir, name),
				}
				m.quitting = true
				return m, tea.Quit
			}
		}
		return m, tea.Batch(cmds...)

	case ProjectListModeDelete:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "y":
				idx := m.list.Index()
				if idx < len(m.projects) {
					if err := m.projectManager.DeleteProject(m.projects[idx].Name); err != nil {
						m.message = "Fehler: " + err.Error()
					} else {
						m.message = "Projekt gelöscht"
						m.projects, _ = m.projectManager.ListProjects()
						items := make([]list.Item, len(m.projects))
						for i, p := range m.projects {
							items[i] = projectItem{info: p}
						}
						m.list.SetItems(items)
					}
				}
				m.mode = ProjectListModeSelect
				return m, nil
			case "n", "esc":
				m.mode = ProjectListModeSelect
				return m, nil
			}
		}
		return m, nil

	default:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "n":
				m.mode = ProjectListModeCreate
				m.nameInput.Reset()
				m.nameInput.Focus()
				return m, textinput.Blink
			case "d":
				if len(m.projects) > 0 {
					m.mode = ProjectListModeDelete
					m.message = "Projekt wirklich löschen? (y/n)"
				}
				return m, nil
			case "enter":
				if len(m.projects) > 0 {
					idx := m.list.Index()
					if idx < len(m.projects) {
						m.selected = &m.projects[idx]
						m.quitting = true
						return m, tea.Quit
					}
				}
			case "q", "ctrl+c":
				m.quitting = true
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m projectListModel) View() string {
	if m.mode == ProjectListModeCreate {
		var b strings.Builder
		b.WriteString(TitleStyle.Render("Neues Projekt") + "\n\n")
		b.WriteString("Name des Projekts:\n")
		b.WriteString(m.nameInput.View() + "\n\n")
		b.WriteString(HelpStyle.Render("Enter: Erstellen | Esc: Abbrechen"))
		return b.String()
	}

	if m.mode == ProjectListModeDelete {
		var b strings.Builder
		b.WriteString(TitleStyle.Render("Projekt löschen") + "\n\n")
		if len(m.projects) > 0 {
			idx := m.list.Index()
			if idx < len(m.projects) {
				b.WriteString(fmt.Sprintf("Projekt '%s' wirklich löschen?\n\n", m.projects[idx].Name))
			}
		}
		b.WriteString(HelpStyle.Render("y: Ja | n: Nein"))
		return b.String()
	}

	var b strings.Builder
	b.WriteString(TitleStyle.Render("Projekt verwalten") + "\n\n")

	if m.message != "" {
		b.WriteString(m.message + "\n\n")
	}

	if len(m.projects) == 0 {
		b.WriteString("Keine Projekte vorhanden.\n\n")
	}

	b.WriteString(m.list.View() + "\n\n")
	b.WriteString(HelpStyle.Render("Enter: Laden | n: Neu | d: Löschen | q: Quit"))
	b.WriteString("\n")

	return b.String()
}

func (m projectListModel) GetSelectedProject() *storage.ProjectInfo {
	return m.selected
}

func (m projectListModel) IsQuitting() bool {
	return m.quitting
}

func (m projectListModel) SetSize(width, height int) {
	m.list.SetSize(width, height-8)
}
