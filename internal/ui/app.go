package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AppState int

const (
	StateProjectView AppState = iota
	StateProjectList
	StateHelp
)

type App struct {
	width       int
	height      int
	state       AppState
	projectName string
	projectView *ProjectView
	projectList *ProjectListView
	helpView    *HelpView
}

func NewApp(projectName string, projectView *ProjectView) *App {
	return &App{
		state:       StateProjectView,
		projectName: projectName,
		projectView: projectView,
		helpView:    NewHelpView(),
	}
}

func (a *App) Init() tea.Cmd {
	return nil
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return a, tea.Quit
		case "ctrl+o":
			if a.state == StateProjectView {
				a.state = StateProjectList
				if a.projectList == nil {
					a.projectList = NewProjectListView()
				}
				return a, nil
			}
		case "?":
			if a.state == StateProjectView {
				a.state = StateHelp
				return a, nil
			}
			if a.state == StateHelp {
				a.state = StateProjectView
				return a, nil
			}
		case "esc":
			if a.state == StateHelp {
				a.state = StateProjectView
				return a, nil
			}
			if a.state == StateProjectList {
				a.state = StateProjectView
				return a, nil
			}
		case "enter":
			if a.state == StateProjectList {
				a.state = StateProjectView
				return a, nil
			}
		}

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
	}

	switch a.state {
	case StateProjectView:
		if a.projectView != nil {
			model, cmd := a.projectView.Update(msg)
			a.projectView = model.(*ProjectView)
			return a, cmd
		}
	case StateProjectList:
		if a.projectList != nil {
			model, cmd := a.projectList.Update(msg)
			a.projectList = model.(*ProjectListView)
			return a, cmd
		}
	case StateHelp:
		if a.helpView != nil {
			model, cmd := a.helpView.Update(msg)
			a.helpView = model.(*HelpView)
			return a, cmd
		}
	}

	return a, nil
}

func (a *App) View() string {
	var content string

	switch a.state {
	case StateProjectView:
		if a.projectView != nil {
			content = a.projectView.View()
		}
	case StateProjectList:
		if a.projectList != nil {
			content = a.projectList.View()
		}
	case StateHelp:
		if a.helpView != nil {
			content = a.helpView.View()
		}
	}

	header := a.renderHeader()
	footer := a.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Top, header, content+"\n", footer)
}

func (a *App) renderHeader() string {
	header := headerStyle.Render("Van Planner - Projekt: " + a.projectName)

	return headerBorderStyle.
		Width(a.width).
		Render(header)
}

func (a *App) renderFooter() string {
	var shortcuts string

	switch a.state {
	case StateProjectView:
		if a.projectView != nil {
			shortcuts = a.projectView.GetShortcuts()
		}
	case StateProjectList:
		if a.projectList != nil {
			shortcuts = a.projectList.GetShortcuts()
		}
	case StateHelp:
		if a.helpView != nil {
			shortcuts = a.helpView.GetShortcuts()
		}
	}

	return footerStyle.
		Width(a.width).
		Render(shortcuts)
}

func (a *App) SetProjectName(name string) {
	a.projectName = name
}

func (a *App) GetCurrentState() AppState {
	return a.state
}

func (a *App) GetProjectView() *ProjectView {
	return a.projectView
}

func (a *App) GetProjectList() *ProjectListView {
	return a.projectList
}

func (a *App) SetProjectList(pl *ProjectListView) {
	a.projectList = pl
}
