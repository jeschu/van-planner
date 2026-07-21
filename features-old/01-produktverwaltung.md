# Feature: Produktverwaltung (CRUD)

## Status
✅ Implementiert

## Beschreibung
Vollständige Verwaltung von Produkten mit Create, Read, Update und Delete Operationen.

## Funktionalität

### Create
- Neues Produkt über `n` Taste erstellen
- Formular mit allen Produkt-Attributen
- Automatische UUID-Generierung
- Aktuelle Kategorie wird übernommen

### Read
- Produkte in gruppierten Liste anzeigen (nach Kategorien)
- Details im Formular bearbeitbar

### Update
- Produkt über `e` Taste bearbeiten
- Alle Attribute änderbar
- Completed-Status bleibt erhalten

### Delete
- Produkt über `d` Taste löschen
- Sofortige Persistenz

## Dateien
- `internal/ui/form.go` – Formular-Implementierung
- `internal/ui/app.go` – CRUD-Logik
- `internal/model/product.go` – Datenmodell

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `n` | Neues Produkt |
| `e` | Produkt bearbeiten |
| `d` | Produkt löschen |

## Datenmodell
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