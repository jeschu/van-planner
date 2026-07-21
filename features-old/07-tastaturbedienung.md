# Feature: Tastaturbedienung

## Status
✅ Implementiert

## Beschreibung
Vollständige Bedienung der Anwendung über Tastatur-Shortcuts mit kontextspezifischer Help-Leiste und detaillierter Hilfe-Seite.

## Funktionalität

### Navigation
| Taste | Aktion |
|-------|--------|
| `j` / `↓` | Nächstes Produkt |
| `k` / `↑` | Vorheriges Produkt |

### Aktionen
| Taste | Aktion |
|-------|--------|
| `Space` | Completed toggle |
| `n` | Neues Produkt |
| `e` | Produkt bearbeiten |
| `d` | Produkt löschen |
| `/` | Suche öffnen |

### Kategorien
| Taste | Aktion |
|-------|--------|
| `K` | Neue Kategorie |
| `E` | Kategorie bearbeiten |
| `D` | Kategorie löschen |

### Projekte
| Taste | Aktion |
|-------|--------|
| `Ctrl+O` | Projekt öffnen |
| `Ctrl+S` | Projekt speichern |

### Formular
| Taste | Aktion |
|-------|--------|
| `Tab` | Nächstes Feld |
| `Shift+Tab` | Vorheriges Feld |
| `Enter` | Speichern (bei letztem Feld) |
| `Esc` | Abbrechen |

### Allgemein
| Taste | Aktion |
|-------|--------|
| `?` | Hilfe-Seite öffnen |
| `q` | Anwendung beenden |
| `Ctrl+C` | Anwendung beenden |

## Dateien
- `internal/ui/app.go` – Globale Shortcuts
- `internal/ui/list.go` – Listen-Shortcuts
- `internal/ui/form.go` – Formular-Shortcuts
- `internal/ui/help.go` – Hilfe-Seite und kontextspezifische Help

## Help-Leiste
Die Help-Leiste zeigt nur die im aktuellen Kontext verfügbaren Shortcuts:

### Listen-Ansicht
```
j/k: Navigation | Space: Toggle | n: Neu | e: Edit | d: Delete | /: Suche | ?: Hilfe | q: Quit
```

### Formular
```
Tab/Enter: Weiter | Shift+Tab: Zurück | Enter: Speichern | Esc: Abbrechen
```

### Hilfe-Seite
Drücke `?` für eine detaillierte Hilfe-Seite mit allen verfügbaren Shortcuts, gruppiert nach Kategorien:
- Navigation
- Aktionen
- Kategorien
- Projekte
- Allgemein