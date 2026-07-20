# Feature: Fortschrittsanzeige

## Status
✅ Implementiert

## Beschreibung
Visuelle Anzeige des Planungsfortschritts mit erledigten Produkten.

## Funktionalität

### Anzeige
- Format: "X/Y erledigt"
- In Status-Leiste unten angezeigt
- Echtzeit-Update bei Toggle

### Completed-Status
- Toggle mit Space-Taste
- Checkbox-Anzeige: `[ ]` / `[✓]`
- Persistenz in JSON

### Visuelle Darstellung
- Erledigte Produkte mit ✓ markiert
- Farbliche Hervorhebung (grün)

## Dateien
- `internal/ui/app.go` – Fortschrittsberechnung
- `internal/ui/list.go` – Checkbox-Rendering
- `internal/ui/styles.go` – CompletedStyle

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `Space` | Completed-Status toggle |

## Datenmodell
```json
{
  "completed": false
}
```

## UI-Beispiel
```
Fortschritt: 3/10 erledigt

[ ] Produkt A
[✓] Produkt B (erledigt)
[ ] Produkt C
```