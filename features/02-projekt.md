# Feature: Projekt

## Status
✅ Implementiert

## Beschreibung
Ein Projekt wird im Ordner `projekte` als JSON-Datei gespeichert. Die Hauptansicht der TUI zeigt die Produkte des Projekts als Tabelle, gruppiert nach Kategorien in der definierten Reihenfolge. Produkte sind nach Index sortiert. Mit den Cursortasten lässt sich ein Produkt auswählen, das im Content-Bereich mit allen Feldern angezeigt wird. Bei breiter Darstellung wird horizontales Scrollen unterstützt.

## Funktionalität

### Projekt-Speicherstruktur
- Projekte werden als JSON-Dateien im Ordner `projekte/` gespeichert
- Dateiname: `<Projektname>.json`
- JSON-Format enthält Kategorien-Reihenfolge und Produktliste

### Produktliste (Tabelle)
- Produkte werden nach Kategorien gruppiert angezeigt
- Kategorien-Reihenfolge entspricht der Definition im `categories`-Array
- Produkte innerhalb einer Kategorie sind nach `index` sortiert
- Tabellarische Darstellung der Produkte
- Cursor-Navigation durch die Liste

### Produkt-Auswahl
- Auswahl eines Produkts mit Cursortasten (↑, ↓)
- Ausgewähltes Produkt wird im Content-Bereich detailliert angezeigt
- View scrollt automatisch, sodass ausgewähltes Produkt möglichst mittig angezeigt wird
- Horizontales Scrollen (←, →) bei breiter Darstellung (wenn Zeile breiter als Terminal)

### Angezeigte Produkt-Felder
- `index` – Laufende Nummer
- `name` – Produktname
- `count` – Anzahl (kann null sein)
- `estimatedCost` – Geschätzte Kosten
- `actualCost` – Tatsächliche Kosten
- `shippingCost` – Versandkosten
- `shopLink` – Shop-Link
- `notes` – Notizen
- `completed` – Abschluss-Status
- `category` – Kategorie-Zuordnung

## Dateien
- `internal/model/project.go` – Projekt-Datenmodell (Project, Product)
- `internal/storage/project.go` – Projekt-Speicher (Laden/Speichern von JSON)
- `internal/ui/project_view.go` – Projekt-Ansicht mit Tabelle und Detail-Anzeige
- `projekte/*.json` – Projekt-Dateien

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `↑` / `k` | Vorheriges Produkt auswählen |
| `↓` / `j` | Nächstes Produkt auswählen |
| `←` | Nach links scrollen (bei breiter Ansicht) |
| `→` | Nach rechts scrollen (bei breiter Ansicht) |
| `Space` | Produkt als erledigt markieren (completed toggle) |

## Datenmodell
```json
{
    "categories": [
        "Fahrzeug",
        "Werkzeug",
        "Dämmung",
        "Heizung",
        "Wasser",
        "Fenster",
        "Elektrik"
    ],
    "products": [
        {
            "index": 1,
            "name": "Ducato",
            "count": null,
            "estimatedCost": 0.0,
            "actualCost": 0.0,
            "shippingCost": 0.0,
            "shopLink": "",
            "notes": "",
            "completed": false,
            "category": "Fahrzeug"
        }
    ]
}
```

## UI-Beispiel
```
┌─────────────────────────────────────────────────────────────┐
│  Van Planner                                    Projekt: …  │
├─────────────────────────────────────────────────────────────┤
│  Fahrzeug                                                   │
│  > [ ] 1. Ducato                                            │
│                                                             │
│  Werkzeug                                                   │
│    [ ] 2. Blindnietzange                                    │
│    [ ] 3. Akku-Schrauber/Bohrer                             │
│                                                             │
│  ─────────────────────────────────────────────────────────  │
│  Ausgewähltes Produkt:                                      │
│  Name: Ducato                                               │
│  Kategorie: Fahrzeug                                        │
│  Geschätzte Kosten: 0,00 €                                  │
│  Tatsächliche Kosten: 0,00 €                                │
│  Versandkosten: 0,00 €                                      │
│  Shop-Link: https://...                                     │
│  Notizen: …                                                 │
│  Status: ☐ offen                                            │
├─────────────────────────────────────────────────────────────┤
│  [↑]↑ [↓]↓ [←][→]Scroll [Space]Erledigt [q]Quit            │
└─────────────────────────────────────────────────────────────┘
```

## Abhängigkeiten
- Feature 01 (TUI-Layout) muss implementiert sein
- Bubble Tea (bereits in go.mod)
- Lipgloss für Layout (bereits in go.mod)
- Bubbles für Tabelle/Viewport (bereits in go.mod)

## Offene Fragen
- Soll die Detail-Anzeige fixiert sein oder als Overlay/Modal?
- Sollen Produkte auch horizontal mit Tab/Shift+Tab navigierbar sein?