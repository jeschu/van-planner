# Feature: Hilfe-Seite

## Status
✅ Implementiert

## Beschreibung
Eine detaillierte Hilfe-Seite zeigt alle verfügbaren Tastatur-Shortcuts, gruppiert nach Kategorien. Die Help-Leiste am unteren Rand zeigt kontextspezifisch nur die aktuell verfügbaren Shortcuts.

## Funktionalität

### Kontextspezifische Help-Leiste
Die Help-Leiste am unteren Bildschirmrand zeigt nur die im aktuellen Modus verfügbaren Shortcuts:

- **Listen-Ansicht**: Navigation, Produkt-Aktionen, Kategorien, Projekte, Suche
- **Formular**: Feld-Navigation, Speichern, Abbrechen
- **Projekt-Liste**: Projekt-Aktionen
- **Kategorie-Formular**: Speichern, Abbrechen

### Hilfe-Seite
Drücke `?` für eine vollständige Hilfe-Seite mit allen Shortcuts:

- **Navigation**: j/k für Navigation, Space für Toggle
- **Aktionen**: n (neu), e (edit), d (delete), / (Suche)
- **Kategorien**: K (neu), E (edit), D (delete)
- **Projekte**: Ctrl+O (öffnen), Ctrl+S (speichern)
- **Allgemein**: ? (Hilfe), q (quit), Ctrl+C (quit)

### Hilfe-Seite schließen
- Beliebige Taste drücken zum Schließen

## Dateien
- `internal/ui/help.go` – Hilfe-Seite und kontextspezifische Help-Funktionen
- `internal/ui/app.go` – Integration der Hilfe in die App

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `?` | Hilfe-Seite öffnen/schließen |

## UI-Beispiel
### Help-Leiste (Listen-Ansicht)
```
j/k: Navigation | Space: Toggle | n: Neu | e: Edit | d: Delete | /: Suche | ?: Hilfe | q: Quit
```

### Hilfe-Seite
```
Hilfe – Alle Tastatur-Shortcuts

Navigation
  j/↓:   Nächstes Produkt
  k/↑:   Vorheriges Produkt
  Space: Completed toggle

Aktionen
  n:   Neues Produkt
  e:   Produkt editieren
  d:   Produkt löschen
  /:   Suche

Kategorien
  K:   Neue Kategorie
  E:   Kategorie editieren
  D:   Kategorie löschen

Projekte
  Ctrl+O: Projekt öffnen
  Ctrl+S: Projekt speichern

Allgemein
  ?:      Hilfe
  q:      Beenden
  Ctrl+C: Beenden

Drücke eine beliebige Taste zum Schließen
```

## Abhängigkeiten
- Feature 07 (Tastaturbedienung) muss implementiert sein