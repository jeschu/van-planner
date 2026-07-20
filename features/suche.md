# Feature: Suche

## Status
✅ Implementiert

## Beschreibung
Volltextsuche über Produkte für schnelles Finden.

## Funktionalität

### Suchfelder
- Produktname
- Notizen
- Kategorie

### Suchverhalten
- Case-insensitive
- Teilstring-Matching
- Echtzeit-Filterung

### Such-UI
- Suchfeld über `/` Taste öffnen
- Eingabe mit Live-Feedback
- Mit Enter bestätigen
- Mit Esc abbrechen/zurücksetzen

## Dateien
- `internal/ui/list.go` – Such-Implementierung
- `internal/ui/app.go` – Such-Modal Integration

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `/` | Suche öffnen |
| `Enter` | Suche ausführen |
| `Esc` | Suche schließen |

## Code-Beispiel
```go
func filterBySearch(data model.Data, query string) model.Data {
    query = strings.ToLower(query)
    for _, p := range data.Products {
        if strings.Contains(strings.ToLower(p.Name), query) ||
           strings.Contains(strings.ToLower(p.Notes), query) ||
           strings.Contains(strings.ToLower(p.Category), query) {
            filtered.Products = append(filtered.Products, p)
        }
    }
    return filtered
}
```