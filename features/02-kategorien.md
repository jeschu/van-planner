# Feature: Kategorien

## Status
✅ Implementiert

## Beschreibung
Organisation von Produkten in Kategorien für bessere Strukturierung.

## Funktionalität

### Standard-Kategorien
1. Elektrik
2. Wasser
3. Küche
4. Schlafen
5. Stauraum

### Kategorie-Anzeige
- Tabs mit Produktanzahl pro Kategorie
- Aktive Kategorie hervorgehoben
- Kategorie-Index (1-5) angezeigt

### Navigation
- Wechsel mit Pfeiltasten (←/→)
- Direkte Auswahl mit Tasten 1-5
- Filterung der Produktliste

### Kategorie bei Create
- Aktuelle Kategorie wird im Formular vorbelegt
- Kategorie im Formular änderbar

## Dateien
- `internal/ui/list.go` – Kategorie-Tabs und Filterung
- `internal/model/product.go` – Kategorie im Produktmodell
- `data/products.json` – Kategorien definierbar

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `←` / `→` | Vorherige/Nächste Kategorie |
| `1` - `5` | Direkte Kategorie-Auswahl |

## Erweiterbarkeit
Kategorien können in `data/products.json` angepasst werden:
```json
{
  "categories": ["Neue Kategorie", "Weitere Kategorie"]
}
```