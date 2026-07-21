package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/jens/van-planner/internal/controller"
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

	controller := controller.NewController(configStorage, projectStorage)

	if err := controller.Initialize(); err != nil {
		log.Fatalf("Fehler beim Initialisieren: %v", err)
	}

	app := controller.CreateApp()

	projects, err := controller.GetProjectList()
	if err == nil && len(projects) > 0 {
		if projectList := ui.NewProjectListView(); projectList != nil {
			projectList.SetProjects(projects)
		}
	}

	p := tea.NewProgram(app, tea.WithAltScreen())

	_, err = p.Run()
	if err != nil {
		fmt.Printf("Fehler beim Ausführen der Anwendung: %v\n", err)
		os.Exit(1)
	}

	if err := controller.SaveCurrentState(); err != nil {
		fmt.Printf("Fehler beim Speichern: %v\n", err)
	}

	absPath, _ := filepath.Abs(projectsDir)
	fmt.Printf("Projekt gespeichert in: %s\n", absPath)
}
