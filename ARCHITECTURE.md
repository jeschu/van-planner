# Van Planner - Architektur

## Übersicht

Van Planner ist eine terminal-basierte Anwendung zur Planung von Campervan-Ausbau-Produkten. Die Anwendung folgt dem **Model-View-Controller (MVC)**-Pattern und nutzt die **Bubble Tea**-Architektur.

## Architektur-Übersicht

```
┌─────────────────────────────────────────────────────────────┐
│                         main.go                              │
│                          Entry Point                         │
└─────────────────────────────────────────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────┐
│                   internal/controller                        │
│  ┌──────────────────────────────────────────────────────┐   │
│  │                  controller.go                       │   │
│  │  - Controller                                        │   │
│  │  - Koordiniert UI, Storage und Business-Logik        │   │
│  │  - Projekt-Management                                │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
                               │
               ┌───────────────┴───────────────┐
               │                               │
               ▼                               ▼
┌──────────────────────────┐    ┌──────────────────────────┐
│   internal/ui            │    │   internal/storage       │
│  ┌────────────────────┐  │    │  ┌────────────────────┐  │
│  │   app.go           │  │    │  │   config.go        │  │
│  │   project_view.go  │  │    │  │   - ConfigStorage  │  │
│  │   project_list.go  │  │    │  │   - Load/Save      │  │
│  │   help.go          │  │    │  └────────────────────┘  │
│  └────────────────────┘  │    │  ┌────────────────────┐  │
│                          │    │  │   project.go       │  │
│                          │    │  │   - ProjectStorage │  │
│                          │    │  │   - Load/Save/List │  │
│                          │    │  └────────────────────┘  │
└──────────────────────────┘    └──────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────┐
│                   internal/model                             │
│  ┌──────────────────────────────────────────────────────┐   │
│  │   project.go                                         │   │
│  │   - Project                                          │   │
│  │   - Product                                          │   │
│  └──────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────┘
```

## Verzeichnisstruktur

```
van-planner/
├── main.go                  # Application Entry Point
├── internal/
│   ├── controller/
│   │   └── controller.go    # Controller: Koordiniert UI & Storage
│   ├── model/
│   │   └── project.go       # Datenmodelle (Project, Product)
│   ├── storage/
│   │   ├── config.go        # Config-Persistenz
│   │   └── project.go       # Projekt-Persistenz
│   └── ui/
│       ├── app.go           # Haupt-Model (Bubble Tea)
│       ├── project_view.go  # Projekt-Ansicht
│       ├── project_list.go  # Projekt-Liste
│       └── help.go          # Hilfe-Ansicht
├── projekte/
│   └── *.json               # Projekt-Daten
├── go.mod
├── go.sum
├── README.md
└── ARCHITECTURE.md
```

## Design-Prinzipien

### 1. Bubble Tea Architecture (Elm Architecture)

Die Anwendung folgt der Elm-Architektur, die von Bubble Tea implementiert wird:

```
┌─────────────┐
│     Msg     │ ──┐
└─────────────┘   │
                  ▼
┌─────────────┐  ┌─────────────┐
│    Model    │◄─┤   Update    │
└─────────────┘  └─────────────┘
       │                ▲
       │                │
       ▼                │
┌─────────────┐         │
│    View     │─────────┘
└─────────────┘
```

- **Model**: Hält den Anwendungszustand (`App`, `listModel`, `formModel`)
- **Update**: Verarbeitet Nachrichten und aktualisiert den Zustand
- **View**: Rendert den aktuellen Zustand als String
- **Msg**: Nachrichten von Events (Tastatur, Timer, etc.)

### 2. Separation of Concerns

Jede Komponente hat eine klare, einzelne Verantwortung:

| Komponente | Verantwortung |
|------------|---------------|
| `controller/` | Orchestrierung, Business-Logik, Projekt-Management |
| `model/` | Datenstrukturen definieren |
| `storage/` | Datenpersistenz (JSON I/O) |
| `ui/` | Präsentation und Interaktion |

### 3. Immutable State Updates

Zustandsänderungen erfolgen durch Erzeugung neuer Modelle statt Mutation:

```go
func (p *ProjectView) toggleProductCompleted(index int) {
    newProducts := make([]model.Product, len(p.project.Products))
    for i, product := range p.project.Products {
        newProducts[i] = product
    }
    newProducts[index].Completed = !newProducts[index].Completed
    p.project.Products = newProducts
}
```

Vorteile:
- Vorhersehbarer State-Flow
- Einfachere Debugging-Möglichkeiten
- Thread-safe durch keine Seiteneffekte

### 4. Dependency Injection

Die Storage-Komponente wird injiziert, was Testbarkeit ermöglicht:

```go
func NewApp(storage *storage.JSONStorage) (*App, error) {
    // Storage wird von außen bereitgestellt
}
```

## Komponenten

### Controller (`internal/controller/controller.go`)

**Verantwortung**: Orchestrierung und Business-Logik

- Initialisiert Anwendung
- Koordiniert UI und Storage
- Projekt-Management (Laden, Speichern, Wechseln)
- Kapselt Datenzugriff von der UI

```go
type Controller struct {
    configStorage  *storage.ConfigStorage
    projectStorage *storage.ProjectStorage
    currentProject *model.Project
    app            *ui.App
}

func (c *Controller) SwitchProject(projectName string) error
func (c *Controller) SaveCurrentState() error
```

### Model (`internal/model/project.go`)

**Verantwortung**: Datenstrukturen definieren

```go
type Project struct {
    Categories []string  `json:"categories"`
    Products   []Product `json:"products"`
}

type Product struct {
    Index         int            `json:"index"`
    Name          string         `json:"name"`
    Count         *int           `json:"count"`
    EstimatedCost float64        `json:"estimatedCost"`
    ActualCost    float64        `json:"actualCost"`
    ShippingCost  float64        `json:"shippingCost"`
    ShopLink      string         `json:"shopLink"`
    Notes         string         `json:"notes"`
    Completed     bool           `json:"completed"`
    Category      string         `json:"category"`
    CustomFields  map[string]any `json:"customFields"`
}
```

### Storage (`internal/storage/`)

**Verantwortung**: Datenpersistenz

- `Load()`: Lädt Daten aus JSON-Datei
- `Save()`: Speichert Daten in JSON-Datei
- Erstellt automatisch Verzeichnisstruktur

### UI (`internal/ui/`)

#### `app.go` – Haupt-Model

- Koordiniert alle UI-Komponenten
- Verwaltet States (ProjectView, ProjectList, Help)
- Delegiert Messages an aktive Views

#### `project_view.go` – Projekt-Ansicht

- Zeigt Produkte nach Kategorien gruppiert
- Toggle Completed-Status (Immutable Updates)
- Horizontales und vertikales Scrollen
- Auto-Scroll zum Cursor

#### `project_list.go` – Projekt-Liste

- Zeigt verfügbare Projekte
- Auswahl mit Navigation
- Viewport-basiertes Rendering

#### `help.go` – Hilfe-Ansicht

- Zeigt Keyboard-Shortcuts
- Dokumentierte Funktionen

## Datenfluss

```
Benutzer-Eingabe (Taste)
        │
        ▼
┌───────────────┐
│   tea.Msg     │
└───────────────┘
        │
        ▼
┌───────────────┐
│  App.Update() │ ──► State-Änderung
└───────────────┘
        │
        ▼
┌───────────────┐
│  Controller   │ ──► Business-Logik
└───────────────┘
        │
        ▼
┌───────────────┐
│  storage.Save()│ ──► JSON-Datei
└───────────────┘
        │
        ▼
┌───────────────┐
│  App.View()   │ ──► Terminal-Output
└───────────────┘
```

## State Management

### App-States (Modi)

```go
type mode int

const (
    modeList mode = iota   // Hauptansicht
    modeCreate            // Formular: Neu
    modeEdit              // Formular: Bearbeiten
    modeDelete            // Lösch-Bestätigung
)
```

### Listen-Status

- Aktuelle Kategorie (Index)
- Suchzustand (aktiv/inaktiv)
- Ausgewählter Index

## Erweiterbarkeit

### Neue Kategorien

Kategorien sind dynamisch in `data/config.json` definierbar.

### Neue Produkt-Attribute

Attribute können im `Product`-Modell ergänzt werden:

```go
type Product struct {
    // ... bestehende Felder
    NewField string `json:"newField"`
}
```

### Neue UI-Komponenten

Neue Views können als separate Modelle implementiert werden:

```go
type newViewModel struct {
    // State
}

func (m newViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd)
func (m newViewModel) View() string
```

## Abhängigkeiten

| Package | Zweck |
|---------|-------|
| `bubbletea` | TUI-Framework (Elm-Architektur) |
| `bubbles` | Vorgefertigte TUI-Komponenten |
| `lipgloss` | Styling für Terminal |
| `uuid` | UUID-Generierung für Produkt-IDs |

## Testbarkeit

Die Architektur ermöglicht Unit-Tests durch:

1. **Pure Functions**: `Update()` und `View()` sind deterministisch
2. **Dependency Injection**: Storage kann gemockt werden
3. **Isolierte Komponenten**: Jedes Modell ist separat testbar

Beispiel:

```go
func TestListUpdate(t *testing.T) {
    model := newListModel(testData)
    updated, _ := model.Update(tea.KeyMsg{Type: tea.KeySpace})
    // Assert...
}
```

## Performance

- **Speicher**: Alle Daten im RAM, JSON nur bei Änderungen
- **Rendering**: Lipgloss optimiert Terminal-Output
- **Navigation**: O(1) für Kategorie-Wechsel durch Index

## Sicherheit

- Keine externen Netzwerkaufrufe
- Lokale Datenspeicherung
- Keine Secrets im Code