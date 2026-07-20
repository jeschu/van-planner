# Feature: Tastaturbedienung

## Status
✅ Implementiert

## Beschreibung
Vollständige Bedienung der Anwendung über Tastatur-Shortcuts.

## Funktionalität

### Navigation
| Taste | Aktion |
|-------|--------|
| `j` / `↓` | Runter navigieren |
| `k` / `↑` | Hoch navigieren |
| `←` / `→` | Kategorie wechseln |
| `1` - `5` | Direkte Kategorie-Auswahl |

### Aktionen
| Taste | Aktion |
|-------|--------|
| `Space` | Completed toggle |
| `n` | Neues Produkt |
| `e` | Produkt bearbeiten |
| `d` | Produkt löschen |
| `/` | Suche öffnen |

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
| `q` | Anwendung beenden |
| `Ctrl+C` | Anwendung beenden |

## Dateien
- `internal/ui/app.go` – Globale Shortcuts
- `internal/ui/list.go` – Listen-Shortcuts
- `internal/ui/form.go` – Formular-Shortcuts

## Help-Leiste
Alle Shortcuts werden in der Help-Leiste unten angezeigt:
```
j/k: Navigation | ←/→: Kategorie | Space: Toggle | n: Neu | e: Edit | d: Delete | /: Suche | q: Quit
```