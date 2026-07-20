# Van Planner - Architektur

## Übersicht

Van Planner ist eine terminal-basierte Anwendung zur Planung von Campervan-Ausbau-Produkten. Die Anwendung folgt dem **Model-View-Controller (MVC)**-Pattern und nutzt die **Bubble Tea**-Architektur.

## Architektur-Übersicht

```
┌─────────────────────────────────────────────────────────────┐
│                      cmd/van-planner                         │
│                         (main.go)                            │
│                          Entry Point                         │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        internal/ui                           │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐    │
│  │  app.go  │  │ list.go  │  │ form.go  │  │ styles.go│    │
│  │  (Model) │  │  (View)  │  │ (View)   │  │ (Styles) │    │
│  └──────────┘  └──────────┘  └──────────┘  └──────────┘    │
└─────────────────────────────────────────────────────────────┘
                              │
              ┌───────────────┴───────────────┐
              │                               │
              ▼                               ▼
┌──────────────────────────┐    ┌──────────────────────────┐
│   internal/model         │    │   internal/storage       │
│  ┌────────────────────┐  │    │  ┌────────────────────┐  │
│  │   product.go       │  │    │  │   json.go          │  │
│  │   - Product        │  │    │  │   - JSONStorage    │  │
│  │   - Data           │  │    │  │   - Load()         │  │
│  └────────────────────┘  │    │  │   - Save()         │  │
│                          │    │  └────────────────────┘  │
└──────────────────────────┘    └──────────────────────────┘
```

## Verzeichnisstruktur

```
van-planner/
├── cmd/
│   └── van-planner/
│       └── main.go          # Application Entry Point
├── internal/
│   ├── model/
│   │   └── product.go       # Datenmodelle (Product, Data)
│   ├── storage/
│   │   └── json.go          # JSON-Persistenz
│   └── ui/
│       ├── app.go           # Haupt-Model (Bubble Tea)
│       ├── list.go          # Listenansicht
│       ├── form.go          # Formular für Create/Edit
│       └── styles.go        # Lipgloss Styles
├── data/
│   └── products.json        # Datenspeicher
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
| `model/` | Datenstrukturen und Business-Logik |
| `storage/` | Datenpersistenz (JSON I/O) |
| `ui/` | Präsentation und Interaktion |

### 3. Immutable State Updates

Zustandsänderungen erfolgen durch Erzeugung neuer Modelle:

```go
func (m listModel) Update(msg tea.Msg) (listModel, tea.Cmd) {
    // Erzeuge neues Modell statt Mutation
    newModel := m
    newModel.data = updatedData
    return newModel, nil
}
```

### 4. Dependency Injection

Die Storage-Komponente wird injiziert, was Testbarkeit ermöglicht:

```go
func NewApp(storage *storage.JSONStorage) (*App, error) {
    // Storage wird von außen bereitgestellt
}
```

## Komponenten

### Model (`internal/model/product.go`)

**Verantwortung**: Datenstrukturen definieren

```go
type Product struct {
    ID           string
    Name         string
    Category     string
    Completed    bool
    Price        float64
    ShopLink     string
    Notes        string
    CustomFields map[string]interface{}
}

type Data struct {
    Categories []string
    Products   []Product
}
```

### Storage (`internal/storage/json.go`)

**Verantwortung**: Datenpersistenz

- `Load()`: Lädt Daten aus JSON-Datei
- `Save()`: Speichert Daten in JSON-Datei
- Erstellt automatisch Verzeichnisstruktur

### UI (`internal/ui/`)

#### `app.go` – Haupt-Model

- Koordiniert alle UI-Komponenten
- Verwaltet Modi (Liste, Create, Edit)
- Persistiert Daten nach Änderungen

#### `list.go` – Listenansicht

- Zeigt Produkte als Liste
- Kategorie-Filterung
- Suchfunktionalität
- Tastatur-Navigation

#### `form.go` – Formular

- Eingabe von Produkt-Attributen
- Validierung
- Navigation zwischen Feldern

#### `styles.go` – Styling

- Zentrale Definition aller Lipgloss-Styles
- Konsistentes Farbschema

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

Kategorien sind dynamisch in `data/products.json` definierbar.

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