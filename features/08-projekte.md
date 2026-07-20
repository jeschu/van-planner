# Feature: Projekte

## Status
🔲 Geplant

## Beschreibung
Laden und Speichern von verschiedenen Projekten als separate JSON-Dateien.

## Funktionalität

### Projekt erstellen
- Neues Projekt über Menü erstellen
- Projektname als Dateiname (z.B. `my-van.json`)
- Leere Projektstruktur mit Standard-Kategorien

### Projekt laden
- Projekt aus `projekte/` Verzeichnis laden
- Projektliste beim Start anzeigen oder über Menü
- Zuletzt verwendetes Projekt merken

### Projekt speichern
- Automatisch im `projekte/` Verzeichnis speichern
- Projektname im Titel anzeigen

### Projekt wechseln
- Über Projekt-Menü zwischen Projekten wechseln
- Aktuelles Projekt in Status-Leiste anzeigen

## Dateien
- `internal/storage/project.go` – Projekt-Management
- `internal/ui/project_list.go` – Projektauswahl
- `internal/ui/app.go` – Projekt-Wechsel Logik
- `projekte/*.json` – Projekt-Dateien

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `Ctrl+O` | Projekt öffnen |
| `Ctrl+S` | Projekt speichern |
| `Ctrl+N` | Neues Projekt |
| `Ctrl+Q` | Projekt schließen |

## Datenmodell
Jedes Projekt ist eine separate JSON-Datei:
```json
{
  "name": "Projektname",
  "categories": ["Elektrik", "Wasser", "Küche", "Schlafen", "Stauraum"],
  "products": [...]
}
```

## Verzeichnisstruktur
```
van-planner/
├── projekte/
│   ├── ausbau-2024.json
│   ├── wohnmobil.json
│   └── ...
```