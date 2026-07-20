package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/jens/van-planner/internal/model"
	"github.com/jens/van-planner/internal/storage"
)

type App struct {
	data     model.Data
	storage  *storage.JSONStorage
	list     listModel
	quitting bool
}

func NewApp(storage *storage.JSONStorage) (*App, error) {
	data, err := storage.Load()
	if err != nil {
		return nil, err
	}

	app := &App{
		data:    data,
		storage: storage,
		list:    newListModel(data),
	}

	return app, nil
}

func (a *App) Init() tea.Cmd {
	return nil
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			a.quitting = true
			return a, tea.Quit
		}

	case tea.WindowSizeMsg:
		a.list.SetSize(msg.Width, msg.Height)
	}

	var cmd tea.Cmd
	a.list, cmd = a.list.Update(msg)

	a.data = a.list.GetData()

	if err := a.storage.Save(a.data); err != nil {
		return a, nil
	}

	return a, cmd
}

func (a *App) View() string {
	if a.quitting {
		return "Bis bald!\n"
	}
	return a.list.View()
}
