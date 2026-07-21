package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/jens/van-planner/internal/model"
	"github.com/jens/van-planner/internal/storage"
	"github.com/jens/van-planner/internal/ui"
)

func main() {
	projectsDir := "projekte"

	if err := os.MkdirAll(projectsDir, 0755); err != nil {
		log.Fatalf("Fehler beim Erstellen des Projekte-Ordners: %v", err)
	}

	configStorage := storage.NewConfigStorage(projectsDir)
	projectStorage := storage.NewProjectStorage(projectsDir)

	config, err := configStorage.Load()
	if err != nil {
		config = &storage.Config{
			DefaultCategories: []string{
				"Fahrzeug", "Werkzeug", "Dämmung", "Heizung",
				"Wasser", "Fenster", "Elektrik",
			},
			Projects:    []string{},
			LastProject: 0,
		}
	}

	var projectName string
	if len(config.Projects) > 0 {
		if config.LastProject >= 0 && config.LastProject < len(config.Projects) {
			projectName = config.Projects[config.LastProject]
		} else {
			projectName = config.Projects[0]
		}
	} else {
		projectName = "Neues Projekt"
	}

	project, err := projectStorage.Load(projectName)
	if err != nil {
		project = &model.Project{
			Categories: config.DefaultCategories,
			Products:   []model.Product{},
		}
	}

	projectView := ui.NewProjectView(project)
	app := ui.NewApp(projectName, projectView)

	projects, err := projectStorage.List()
	if err == nil && len(projects) > 0 {
		if projectList := ui.NewProjectListView(); projectList != nil {
			projectList.SetProjects(projects)
		}
	}

	p := tea.NewProgram(app, tea.WithAltScreen())

	model, err := p.Run()
	if err != nil {
		fmt.Printf("Fehler beim Ausführen der Anwendung: %v\n", err)
		os.Exit(1)
	}

	if appModel, ok := model.(*ui.App); ok {
		if appModel.GetCurrentState() == ui.StateProjectView {
			if pv := appModel.GetProjectView(); pv != nil {
				if err := projectStorage.Save(projectName, pv.GetProject()); err != nil {
					fmt.Printf("Fehler beim Speichern des Projekts: %v\n", err)
				}
			}
		}
	}

	absPath, _ := filepath.Abs(projectsDir)
	fmt.Printf("Projekt gespeichert in: %s\n", absPath)
}
