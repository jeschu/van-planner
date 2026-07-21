# Feature: Gruppierte Produktliste

## Status
✅ Implementiert

## Beschreibung
Alle Produkte werden in einer tabellarischen Ansicht angezeigt, gruppiert nach Kategorien. Jede Zeile zeigt Produktname, geschätzte Kosten, tatsächliche Kosten und Link. Kategorie-Summen und Gesamtsumme werden automatisch berechnet.

## Funktionalität

### Tabellarische Anzeige
- Alle Produkte werden in einer Tabelle angezeigt
- Spalten: Produktname, Kosten geschätzt, Kosten tatsächlich, Link
- Produkte sind nach Kategorien gruppiert
- Kategorie-Header zeigen den Kategorie-Namen
- Zwischensumme pro Kategorie (geschätzt und tatsächlich)
- Gesamtsumme aller Kategorien am Ende der Tabelle
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
- `internal/model/product.go` – Datenmodell (EstimatedCost, ActualCost)
- `internal/ui/list.go` – Tabellarische Ansicht mit Summenberechnung
- `internal/ui/form.go` – Formular mit zwei Kosten-Feldern

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `j` / `↓` | Nächstes Produkt |
| `k` / `↑` | Vorheriges Produkt |
| `Space` | Completed-Status toggle |
| `/` | Suche öffnen |

## Datenmodell
```json
{
  "id": "uuid",
  "name": "Produktname",
  "category": "Kategorie",
  "completed": false,
  "estimatedCost": 0.0,
  "actualCost": 0.0,
  "shopLink": "",
  "notes": "",
  "customFields": {}
}
```

## UI-Beispiel
```
Van Planner – Planer für deinen Campervan-Ausbau
Projekt: default

Fahrzeug
  Ducato                                          20.000,00 €              0,00 €
  Summe                                           20.000,00 €              0,00 €

Werkzeug
  [ ] Blindnietzange                                 23,00 €              0,00 €  https://amazon.de/...
  [✓] Akku-Schrauber/Bohrer                          30,95 €              0,00 €  https://amazon.de/...
  Summe                                             150,60 €              0,00 €

Gesamtsumme                                       25.025,47 €              0,00 €
```

## Abhängigkeiten
- Feature 01 (Produktverwaltung) muss implementiert sein