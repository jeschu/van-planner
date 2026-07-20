# Feature: Kategorien bearbeiten

## Status
🔲 Geplant

## Beschreibung
Erstellen, Bearbeiten und Löschen von benutzerdefinierten Kategorien.

## Funktionalität

### Kategorie erstellen
- Neue Kategorie über Tastatur-Shortcut hinzufügen
- Eingabe des Kategoriennamens
- Automatische Speicherung in JSON

### Kategorie bearbeiten
- Bestehende Kategorie umbenennen
- Alle Produkte der Kategorie werden aktualisiert

### Kategorie löschen
- Leere Kategorien löschen
- Kategorien mit Produkten nur nach Bestätigung
- Option: Produkte in andere Kategorie verschieben

### Standard-Kategorien
- Vordefinierte Kategorien: Elektrik, Wasser, Küche, Schlafen, Stauraum
- Benutzerdefinierte Kategorien ergänzend

## Dateien
- `internal/ui/category_form.go` – Kategorie-Formular
- `internal/ui/app.go` – Kategorie-CRUD-Logik
- `data/products.json` – Kategorien gespeichert

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `K` | Neue Kategorie erstellen |
| `E` | Kategorie bearbeiten |
| `D` | Kategorie löschen |

## Datenmodell
```json
{
  "categories": ["Elektrik", "Wasser", "Küche", "Schlafen", "Stauraum", "Eigene Kategorie"]
}
```