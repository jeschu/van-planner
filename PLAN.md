# Van Planner - Implementierungsplan

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
  "categories": ["Fahrzeug", "Elektrik", "Wasser", "Küche", "Schlafen", "Stauraum"],
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
├── projekte/
│   └── config.json
├── go.mod
└── go.sum
```

## UI-Komponenten

### Hauptansicht (Liste)
- Kategorien als vertikale Liste
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

## Implementierungsstatus

### ✅ Abgeschlossene Features

| Feature | Status | Beschreibung |
|---------|--------|--------------|
| Produktverwaltung (CRUD) | ✅ Implementiert | Create, Read, Update, Delete mit Formular |
| Kategorien | ✅ Implementiert | 6 Standard-Kategorien mit Tabs und Filterung |
| Kategorien bearbeiten | ✅ Implementiert | Kategorien erstellen, bearbeiten, löschen |
| Fortschrittsanzeige | ✅ Implementiert | X/Y erledigt mit Checkbox-Toggle |
| Suche | ✅ Implementiert | Volltextsuche über Name, Notizen, Kategorie |
| JSON-Persistenz | ✅ Implementiert | Save-on-Change, Load-on-Start |
| Tastaturbedienung | ✅ Implementiert | Vollständige Keyboard-Navigation |
| Projekte | ✅ Implementiert | Projekte laden, speichern, wechseln |

### 🔲 Geplante Features

| Feature | Status | Beschreibung |
|---------|--------|--------------|
| Benutzerdefinierte Felder | 🔲 Geplant | Erweitern des Datenmodells um Custom Fields |
| Export/Import | 🔲 Geplant | CSV/JSON Export und Import |
| Vorschau/Statistiken | 🔲 Geplant | Übersicht über Budget und Fortschritt |

## Abhängigkeiten

```bash
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/bubbles
go get github.com/charmbracelet/lipgloss
```

## Nächste Schritte

1. Feature: Benutzerdefinierte Felder – Feature-Datei erstellen und implementieren
2. Feature: Export/Import – Feature-Datei erstellen und implementieren
3. Feature: Vorschau/Statistiken – Feature-Datei erstellen und implementieren