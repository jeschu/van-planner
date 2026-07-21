# Feature: TUI-Layout

## Status
🔲 Geplant / ⬇️ In Arbeit / ✅ Implementiert

## Beschreibung
Zentrales TUI-Layout mit Header, scrollbarem Content-Bereich und kontext-sensitivem Footer. Jede View liefert ihre eigene Tastaturnavigation über eine `GetShortcuts()`-Methode.

## Funktionalität

### Header
- Zeigt den Titel der App ("Van Planner")
- Zeigt das aktuell gewählte Projekt an
- Wird zentral im Layout gerendert

### Content-Bereich
- Scrollbar bei Bedarf (zu viele Listen-Einträge oder lange Formulare)
- Enthält die aktuelle View (Liste, Formular, Hilfe, etc.)
- Jede View ist ein eigenes Bubble Tea Model

### Footer
- Zeigt kontext-sensitive Tastaturnavigation
- Jede View implementiert `GetShortcuts() string`
- Format: `[Taste]Aktion [Taste]Aktion ...`
- Wird zentral im Layout gerendert

### Layout-Struktur
```
┌─────────────────────────────────────────────┐
│  Van Planner                                │
│  Projekt: Camper Van                        │
├─────────────────────────────────────────────┤
│  [Content - scrollbar bei Bedarf]           │
│  - Produktliste                             │
│  - Formular                                 │
│  - Hilfe                                    │
├─────────────────────────────────────────────┤
│  [j]↓ [k]↑ [n]Neu [e]Edit [d]Delete [q]Quit │
└─────────────────────────────────────────────┘
```

## Dateien
- `internal/ui/layout.go` – Zentrales Layout (Header/Footer/Content)
- `internal/ui/list.go` – Listenansicht mit `GetShortcuts()`
- `internal/ui/form.go` – Formular mit `GetShortcuts()`
- `internal/ui/help.go` – Hilfe-Seite mit `GetShortcuts()`
- `internal/ui/app.go` – Haupt-Model (koordiniert Views)

## Tastatur-Shortcuts

### Listenansicht
| Taste | Aktion |
|-------|--------|
| `j` / `↓` | Nach unten navigieren |
| `k` / `↑` | Nach oben navigieren |
| `←` / `→` | Kategorie wechseln |
| `Space` | Produkt als erledigt markieren |
| `n` | Neues Produkt erstellen |
| `e` | Produkt bearbeiten |
| `d` | Produkt löschen |
| `/` | Suche öffnen |
| `q` | Anwendung beenden |

### Formular
| Taste | Aktion |
|-------|--------|
| `Tab` / `Enter` | Nächstes Feld |
| `Shift+Tab` | Vorheriges Feld |
| `Enter` (letztes Feld) | Speichern |
| `Esc` | Abbrechen |

### Hilfe
| Taste | Aktion |
|-------|--------|
| `Esc` / `?` | Hilfe schließen |

## Datenmodell

Keine Änderungen am Datenmodell erforderlich.

## UI-Beispiel

```
┌─────────────────────────────────────────────┐
│  Van Planner                                │
│  Projekt: Camper Van                        │
├─────────────────────────────────────────────┤
│  Fahrzeug                                   │
│  ☐ Reifen                                   │
│  ☑ Batteriewechsel                          │
│  ☐ Ölwechsel                                │
│                                             │
│  Elektrik                                   │
│  ☐ Solarpanel                               │
│  ☑ Ladegerät                                │
├─────────────────────────────────────────────┘
│  [j]↓ [k]↑ [←][→]Kat [n]Neu [e]Edit [q]Quit │
└─────────────────────────────────────────────┘
```

## Abhängigkeiten
- Bubble Tea (bereits in go.mod)
- Lipgloss für Layout (bereits in go.mod)

## Offene Fragen
- Keine