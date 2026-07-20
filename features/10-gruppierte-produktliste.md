# Feature: Gruppierte Produktliste

## Status
✅ Implementiert

## Beschreibung
Alle Produkte werden in einer einzigen Liste angezeigt, gruppiert nach Kategorien. Kategorie-Header strukturieren die Liste visuell.

## Funktionalität

### Gruppierte Anzeige
- Alle Produkte werden in einer Liste angezeigt
- Produkte sind nach Kategorien gruppiert
- Kategorie-Header zeigen den Kategorie-Namen
- Produkte innerhalb einer Kategorie sortiert

### Navigation
- Navigation mit `j`/`k` oder Pfeiltasten durch gesamte Liste
- Space toggelt Completed-Status des ausgewählten Produkts
- Keine separate Kategorie-Auswahl mehr nötig

### Suche
- Suche mit `/` filtert Produkte über alle Kategorien
- Kategorie-Header werden in der Suche ausgeblendet
- Suche filtert nach Produktname, Notizen und Kategorie

## Dateien
- `internal/ui/list.go` – Gruppierte Listen-Ansicht (createGroupedListItems, groupedListItem)

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `j` / `↓` | Nächstes Produkt |
| `k` / `↑` | Vorheriges Produkt |
| `Space` | Completed-Status toggle |
| `/` | Suche öffnen |

## Datenmodell
Keine Änderungen am Datenmodell. Gruppierung erfolgt nur in der UI-Darstellung.

## UI-Beispiel
```
Van Planner
Planer für deinen Campervan-Ausbau
Projekt: default

Elektrik
  [ ] Solarpanel 400W
  [✓] Laderegler 30A
  [ ] Batterie 200Ah

Wasser
  [ ] Frischwassertank 100L
  [✓] Wasserpumpe
  [ ] Warmwasserbereiter
```

## Abhängigkeiten
- Feature 01 (Produktverwaltung) muss implementiert sein