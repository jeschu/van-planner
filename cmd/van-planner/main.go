package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jens/van-planner/internal/storage"
	"github.com/jens/van-planner/internal/ui"
)

func main() {
	storage := storage.NewJSONStorage("projekte/products.json")

	app, err := ui.NewApp(storage)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim Starten der App: %v\n", err)
		os.Exit(1)
	}

	p := tea.NewProgram(app, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Fehler beim Ausführen der App: %v\n", err)
		os.Exit(1)
	}
}
