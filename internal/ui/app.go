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
	modeCreateCategory
	modeEditCategory
	modeDeleteCategory
	modeProjectList
)

type App struct {
	data           model.Data
	storage        *storage.JSONStorage
	projectManager *storage.ProjectManager
	list           listModel
	form           formModel
	categoryForm   categoryFormModel
	projectList    projectListModel
	currentProject string
	mode           mode
	quitting       bool
	message        string
}

func NewApp(s *storage.JSONStorage) (*App, error) {
	data, err := s.Load()
	if err != nil {
		return nil, err
	}

	pm := storage.NewProjectManager()

	app := &App{
		data:           data,
		storage:        s,
		projectManager: pm,
		list:           newListModel(data),
		mode:           modeList,
		currentProject: "default",
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
			if a.mode == modeList || a.mode == modeProjectList {
				a.quitting = true
				return a, tea.Quit
			}
		case "ctrl+o":
			if a.mode == modeList {
				a.projectList = newProjectListModel(a.projectManager)
				a.mode = modeProjectList
				return a, nil
			}
		case "ctrl+s":
			if a.mode == modeList {
				if err := a.projectManager.SaveProject(a.currentProject, a.data); err != nil {
					a.message = "Fehler beim Speichern: " + err.Error()
				} else {
					a.message = "Projekt gespeichert"
				}
				return a, nil
			}
		case "ctrl+n":
			if a.mode == modeList {
				a.projectList = newProjectListModel(a.projectManager)
				a.mode = modeProjectList
				return a, nil
			}
		case "ctrl+q":
			if a.mode == modeList {
				a.projectList = newProjectListModel(a.projectManager)
				a.mode = modeProjectList
				return a, nil
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
		case "K":
			if a.mode == modeList {
				a.categoryForm = newCategoryFormModel("create", "")
				a.mode = modeCreateCategory
				return a, textinput.Blink
			}
		case "E":
			if a.mode == modeList {
				currentCat := a.list.GetCurrentCategory()
				a.categoryForm = newCategoryFormModel("edit", currentCat)
				a.mode = modeEditCategory
				return a, textinput.Blink
			}
		case "D":
			if a.mode == modeList {
				currentCat := a.list.GetCurrentCategory()
				hasProducts := false
				for _, p := range a.data.Products {
					if p.Category == currentCat {
						hasProducts = true
						break
					}
				}
				if hasProducts {
					a.message = "Kategorie enthält Produkte. Erst Produkte löschen oder verschieben."
				} else {
					a.deleteCategory(currentCat)
					a.message = "Kategorie gelöscht"
				}
				return a, nil
			}
		}

	case tea.WindowSizeMsg:
		a.list.SetSize(msg.Width, msg.Height)
	}

	if a.mode == modeCreateCategory || a.mode == modeEditCategory {
		var cmd tea.Cmd
		a.categoryForm, cmd = a.categoryForm.Update(msg)
		if a.categoryForm.submitting {
			if a.mode == modeCreateCategory {
				a.createCategory()
			} else if a.mode == modeEditCategory {
				a.editCategory()
			}
			a.mode = modeList
			return a, nil
		}
		return a, cmd
	}

	if a.mode == modeProjectList {
		var cmd tea.Cmd
		a.projectList, cmd = a.projectList.Update(msg)
		if a.projectList.IsQuitting() {
			selected := a.projectList.GetSelectedProject()
			if selected != nil {
				a.loadProject(selected.Name)
			}
			a.mode = modeList
			return a, nil
		}
		return a, cmd
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
		a.saveLastProduct()
		return a, cmd
	}
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
	case modeCreateCategory:
		return TitleStyle.Render("Neue Kategorie") + "\n\n" + a.categoryForm.View()
	case modeEditCategory:
		return TitleStyle.Render("Kategorie bearbeiten") + "\n\n" + a.categoryForm.View()
	case modeProjectList:
		return a.projectList.View()
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
		fmt.Sprintf(" %d/%d erledigt | Projekt: %s ", completed, len(a.data.Products), a.currentProject),
	)

	help := HelpStyle.Render("j/k: Navigation | ↑/↓: Kategorie | Space: Toggle | n: Neu | e: Edit | d: Delete | K: Kat.Neu | E: Kat.Edit | D: Kat.Delete | Ctrl+O: Öffnen | Ctrl+S: Speichern | /: Suche | q: Quit")

	if a.message != "" {
		return status + "\n\n" + a.message + "\n\n" + help
	}

	return status + "\n" + help
}

func (a *App) createCategory() {
	name := a.categoryForm.GetCategoryName()
	if name == "" {
		a.message = "Fehler: Kategorienname ist erforderlich"
		return
	}

	for _, cat := range a.data.Categories {
		if cat == name {
			a.message = "Fehler: Kategorie existiert bereits"
			return
		}
	}

	a.data.Categories = append(a.data.Categories, name)
	a.list = newListModel(a.data)
	a.message = "Kategorie erstellt"
	if err := a.storage.Save(a.data); err != nil {
		a.message = "Fehler beim Speichern: " + err.Error()
	}
}

func (a *App) editCategory() {
	newName := a.categoryForm.GetCategoryName()
	if newName == "" {
		a.message = "Fehler: Kategorienname ist erforderlich"
		return
	}

	oldName := a.list.GetCurrentCategory()
	if oldName == newName {
		a.message = "Keine Änderung"
		return
	}

	for _, cat := range a.data.Categories {
		if cat == newName {
			a.message = "Fehler: Kategorie existiert bereits"
			return
		}
	}

	for i, cat := range a.data.Categories {
		if cat == oldName {
			a.data.Categories[i] = newName
			break
		}
	}

	for i := range a.data.Products {
		if a.data.Products[i].Category == oldName {
			a.data.Products[i].Category = newName
		}
	}

	a.list = newListModel(a.data)
	a.message = "Kategorie aktualisiert"
	if err := a.storage.Save(a.data); err != nil {
		a.message = "Fehler beim Speichern: " + err.Error()
	}
}

func (a *App) deleteCategory(name string) {
	for i, cat := range a.data.Categories {
		if cat == name {
			a.data.Categories = append(a.data.Categories[:i], a.data.Categories[i+1:]...)
			break
		}
	}

	a.list = newListModel(a.data)
	if err := a.storage.Save(a.data); err != nil {
		a.message = "Fehler beim Speichern: " + err.Error()
	}
}

func (a *App) loadProject(name string) {
	data, err := a.projectManager.LoadProject(name)
	if err != nil {
		a.message = "Fehler beim Laden: " + err.Error()
		return
	}

	config, err := a.projectManager.LoadProjectConfig(name)
	if err != nil {
		a.message = "Fehler beim Laden der Config: " + err.Error()
		return
	}

	a.data = data
	a.currentProject = name
	a.list = newListModel(data)
	a.storage = storage.NewJSONStorage(fmt.Sprintf("%s/%s.json", storage.ProjectsDir, name))
	a.message = "Projekt geladen: " + name

	if config.LastProductID != "" {
		for i, product := range a.data.Products {
			if product.ID == config.LastProductID {
				a.list.list.Select(i)
				break
			}
		}
	}
}

func (a *App) saveLastProduct() {
	if len(a.data.Products) == 0 {
		return
	}

	idx := a.list.list.Index()
	if idx >= len(a.data.Products) {
		return
	}

	config := model.ProjectConfig{
		LastProductID: a.data.Products[idx].ID,
	}

	if err := a.projectManager.SaveProjectConfig(a.currentProject, config); err != nil {
		a.message = "Fehler beim Speichern der Config: " + err.Error()
	}
}
