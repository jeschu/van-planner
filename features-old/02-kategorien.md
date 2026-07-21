# Feature: Kategorien

## Status
✅ Implementiert

## Beschreibung
Organisation von Produkten in Kategorien für bessere Strukturierung.

## Funktionalität

### Standard-Kategorien
1. Fahrzeug
2. Elektrik
3. Wasser
4. Küche
5. Schlafen
6. Stauraum

### Kategorie-Anzeige
- Kategorien als vertikale Liste links
- Produktanzahl pro Kategorie angezeigt
- Aktive Kategorie hervorgehoben
- Kategorie-Index (1-5) angezeigt

### Navigation
- Wechsel mit Pfeiltasten (↑/↓)
- Direkte Auswahl mit Tasten 1-6
- Filterung der Produktliste

### Kategorie bei Create
- Aktuelle Kategorie wird im Formular vorbelegt
- Kategorie im Formular änderbar

## Dateien
- `internal/ui/list.go` – Kategorie-Tabs und Filterung
- `internal/model/product.go` – Kategorie im Produktmodell
- `projekte/*.json` – Kategorien definierbar

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `↑` / `↓` (im leeren Bereich) | Vorherige/Nächste Kategorie |
| `1` - `6` | Direkte Kategorie-Auswahl |

## Erweiterbarkeit
Kategorien können in `projekte/*.json` angepasst werden:
```json
{
  "categories": ["Neue Kategorie", "Weitere Kategorie"]
}
```