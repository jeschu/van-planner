# Feature: Projekte laden und speichern

## Status
🔲 Geplant

## Beschreibung
Ermöglicht das Laden von Projekten über eine Projektliste und automatisches Speichern beim Beenden. Beim Neustart wird das zuletzt verwendete Projekt automatisch geladen.

## Funktionalität

### Projekt laden (CTRL-O)
- Öffnet eine scrollbare Liste aller Projekte im `projekte/` Ordner
- Projekte werden zeilenweise im Content-Bereich angezeigt
- Auswahl mit Cursor-Tasten (↑/↓)
- Bestätigung der Auswahl mit Enter

### Automatisches Speichern beim Beenden
- Aktuelles Projekt wird beim Beenden der Anwendung automatisch gespeichert

### Automatisches Laden beim Start
- Beim Neustart wird das letzte Projekt aus `projekte/config.json` automatisch geladen

### Konfiguration
- `projekte/config.json` speichert:
  - `defaultCategories`: Standard-Kategorien (wird in späterem Feature verwendet)
  - `projects`: Liste aller Projekt-Namen
  - `lastProject`: Index des zuletzt verwendeten Projekts

## Dateien
- `main.go` – Steuerung des Anwendungs-Lifecycles (Start/Beenden)
- `ui/project_list.go` – Darstellung der Projektliste
- `storage/config.go` – Lesen/Schreiben der config.json
- `projekte/config.json` – Konfigurationsdatei

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `CTRL-O` | Projektliste öffnen |
| `↑` | Vorheriges Projekt in Liste |
| `↓` | Nächstes Projekt in Liste |
| `Enter` | Projekt auswählen und laden |
| `ESC` | Projektliste schließen |

## Datenmodell
```json
{
  "defaultCategories": ["Fahrzeug", "Werkzeug", "Dämmung", "Heizung", "Wasser", "Fenster", "Elektrik"],
  "projects": ["Camper Van"],
  "lastProject": 0
}
```

## UI-Beispiel
```
┌─────────────────────────────────────┐
│ Projekt laden                       │
├─────────────────────────────────────┤
│ > Camper Van                        │
│   Garage Umbau                      │
│   Werkstatt                         │
│   Bootsanbau                        │
│                                     │
│                                     │
└─────────────────────────────────────┘
```

## Abhängigkeiten
- Feature 01 (TUI) muss implementiert sein
- Feature 02 (Projekt) muss implementiert sein
- tview für scrollbare Listen

## Offene Fragen
- Sollen Projekte aus Unterverzeichnissen von `projekte/` geladen werden?
- Soll ein leeres Projekt erstellt werden, wenn kein Projekt existiert?