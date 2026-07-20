package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbletea"
	"github.com/jens/van-planner/internal/model"
	"github.com/jens/van-planner/internal/storage"
)

type mode int

const (
	modeList mode = iota
	modeCreate
	modeEdit
	modeDelete
)

type App struct {
	data     model.Data
	storage  *storage.JSONStorage
	list     listModel
	form     formModel
	mode     mode
	quitting bool
	message  string
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
		mode:    modeList,
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
			if a.mode == modeList {
				a.quitting = true
				return a, tea.Quit
			}
		case "n":
			if a.mode == modeList {
				a.form = newFormModel(a.data, nil, a.list.GetCurrentCategory())
				a.mode = modeCreate
				return a, textinput.Blink
			}
		case "e":
			if a.mode == modeList && len(a.data.Products) > 0 {
				idx := a.list.list.Index()
				if idx < len(a.data.Products) {
					a.form = newFormModel(a.data, &a.data.Products[idx], "")
					a.mode = modeEdit
					return a, textinput.Blink
				}
			}
		case "d":
			if a.mode == modeList && len(a.data.Products) > 0 {
				idx := a.list.list.Index()
				if idx < len(a.data.Products) {
					a.data.Products = append(a.data.Products[:idx], a.data.Products[idx+1:]...)
					a.list = newListModel(a.data)
					a.message = "Produkt gelöscht"
					if err := a.storage.Save(a.data); err != nil {
						a.message = "Fehler beim Speichern: " + err.Error()
					}
					return a, nil
				}
			}
		}

	case tea.WindowSizeMsg:
		a.list.SetSize(msg.Width, msg.Height)
	}

	var cmd tea.Cmd

	switch a.mode {
	case modeCreate:
		a.form, cmd = a.form.Update(msg)
		if a.form.submitting {
			product, err := a.form.GetProduct(a.data)
			if err != nil {
				a.message = "Fehler: " + err.Error()
			} else {
				a.data.Products = append(a.data.Products, product)
				a.message = "Produkt erstellt"
			}
			a.list = newListModel(a.data)
			a.mode = modeList
			if err := a.storage.Save(a.data); err != nil {
				a.message = "Fehler beim Speichern: " + err.Error()
			}
			return a, nil
		}
		return a, cmd

	case modeEdit:
		a.form, cmd = a.form.Update(msg)
		if a.form.submitting {
			product, err := a.form.GetProduct(a.data)
			if err != nil {
				a.message = "Fehler: " + err.Error()
			} else {
				idx := a.list.list.Index()
				if idx < len(a.data.Products) {
					product.ID = a.data.Products[idx].ID
					product.Completed = a.data.Products[idx].Completed
					a.data.Products[idx] = product
					a.message = "Produkt aktualisiert"
				}
			}
			a.list = newListModel(a.data)
			a.mode = modeList
			if err := a.storage.Save(a.data); err != nil {
				a.message = "Fehler beim Speichern: " + err.Error()
			}
			return a, nil
		}
		return a, cmd

	default:
		a.list, cmd = a.list.Update(msg)
		a.data = a.list.GetData()
		if err := a.storage.Save(a.data); err != nil {
			a.message = "Fehler beim Speichern: " + err.Error()
		}
		return a, cmd
	}

	return a, cmd
}

func (a *App) View() string {
	if a.quitting {
		return "Bis bald!\n"
	}

	switch a.mode {
	case modeCreate:
		return TitleStyle.Render("Neues Produkt") + "\n\n" + a.form.View()
	case modeEdit:
		return TitleStyle.Render("Produkt bearbeiten") + "\n\n" + a.form.View()
	default:
		return a.list.View() + a.statusBar()
	}
}

func (a *App) statusBar() string {
	completed := 0
	for _, p := range a.data.Products {
		if p.Completed {
			completed++
		}
	}

	status := StatusBarStyle.Render(
		fmt.Sprintf(" %d/%d erledigt ", completed, len(a.data.Products)),
	)

	help := HelpStyle.Render("j/k: Navigation | ↑/↓: Kategorie | Space: Toggle | n: Neu | e: Edit | d: Delete | /: Suche | q: Quit")

	if a.message != "" {
		return status + "\n\n" + a.message + "\n\n" + help
	}

	return status + "\n" + help
}
