# Van Planner TUI - Implementierungsplan

## Übersicht

Eine TUI-Anwendung in Go zur Planung von Campervan-Ausbauprodukten mit JSON-Datenhaltung.

## Technische Entscheidungen

| Komponente | Auswahl |
|------------|---------|
| TUI-Bibliothek | Bubble Tea (charm.sh) |
| Datenhaltung | JSON-Datei |
| CRUD-Operationen | Vollständig (Create, Read, Update, Delete) |
| Status-Funktion | Checkbox für "erledigt"-Markierung |
| Produkt-Attribute | Benutzerdefiniert erweiterbar |

## Datenmodell

### Produkt-Struktur

```json
{
  "id": "uuid",
  "name": "string",
  "category": "string",
  "completed": false,
  "price": 0.0,
  "shopLink": "string",
  "notes": "string",
  "customFields": {}
}
```

### JSON-Datei-Struktur

```json
{
  "categories": ["Elektrik", "Wasser", "Küche", "Schlafen", "Stauraum"],
  "products": [...]
}
```

## Architektur

### Verzeichnisstruktur

```
van-planner/
├── cmd/
│   └── van-planner/
│       └── main.go
├── internal/
│   ├── model/
│   │   └── product.go
│   ├── storage/
│   │   └── json.go
│   └── ui/
│       ├── app.go
│       ├── list.go
│       ├── form.go
│       └── styles.go
├── data/
│   └── products.json
├── go.mod
└── go.sum
```

## UI-Komponenten

### Hauptansicht (Liste)
- Kategorien als Tabs/Sections
- Produkte pro Kategorie mit Checkbox
- Tastatur-Shortcuts: `j/k` (Navigation), `Space` (Toggle), `n` (Neu), `e` (Edit), `d` (Delete), `q` (Quit)

### Formular (Create/Edit)
- Eingabefelder für Produkt-Attribute
- Kategorie-Auswahl
- Speichern/Abbrechen

### Status-Leiste
- Aktuelle Kategorie
- Fortschrittsanzeige (erledigt/gesamt)
- Menü-Hints

## Implementierungs-Phasen

### Phase 1: Grundgerüst
1. `go mod init` und Bubble Tea installieren
2. Datenmodelle definieren
3. JSON-Storage (Lesen/Schreiben)
4. Minimale TUI-Liste

### Phase 2: CRUD-Operationen
1. Produkt erstellen (Formular)
2. Produkt bearbeiten
3. Produkt löschen
4. JSON persistieren

### Phase 3: UX-Verbesserungen
1. Checkbox-Toggle
2. Kategorie-Filter/Navigation
3. Suche
4. Validierung

### Phase 4: Erweiterte Features
1. Benutzerdefinierte Felder
2. Export/Import
3. Vorschau/Statistiken

## Abhängigkeiten

```bash
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/bubbles
go get github.com/charmbracelet/lipgloss
```

## Nächste Schritte

1. User-Feedback zu diesem Plan
2. Phase 1 implementieren
3. Inkrementell erweitern