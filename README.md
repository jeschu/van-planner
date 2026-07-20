# Van Planner TUI

Eine Terminal-basierte Anwendung in Go zur Planung von Campervan-Ausbau-Produkten.

## Features

- **Produktverwaltung** – Erstellen, Bearbeiten, Löschen von Produkten
- **Kategorien** – Organisation nach Bereichen (Elektrik, Wasser, Küche, etc.)
- **Fortschrittsanzeige** – Behalte den Überblick über erledigte Produkte
- **Suche** – Schnelles Finden von Produkten
- **JSON-Persistenz** – Alle Daten werden lokal gespeichert
- **Tastaturbedienung** – Effiziente Bedienung ohne Maus

## Installation

### Voraussetzungen

- Go 1.21 oder höher

### Bauen

```bash
go build -o van-planner ./cmd/van-planner
```

### Ausführen

```bash
./van-planner
```

Die Daten werden standardmäßig in `data/products.json` gespeichert.

## Bedienung

### Hauptansicht

| Taste | Aktion |
|-------|--------|
| `j` / `↓` | Nach unten navigieren |
| `k` / `↑` | Nach oben navigieren |
| `←` / `→` | Kategorie wechseln |
| `1` - `5` | Direkte Kategorie-Auswahl |
| `Space` | Produkt als erledigt markieren |
| `n` | Neues Produkt erstellen |
| `e` | Produkt bearbeiten |
| `d` | Produkt löschen |
| `/` | Suche öffnen |
| `q` | Anwendung beenden |

### Formular (Create/Edit)

| Taste | Aktion |
|-------|--------|
| `Tab` / `Enter` | Nächstes Feld |
| `Shift+Tab` | Vorheriges Feld |
| `Enter` (bei letztem Feld) | Speichern |
| `Esc` | Abbrechen |

### Suche

| Taste | Aktion |
|-------|--------|
| `/` | Suche öffnen |
| `Enter` | Suche ausführen |
| `Esc` | Suche schließen |

## Produkt-Attribute

Beim Erstellen oder Bearbeiten eines Produkts können folgende Attribute angegeben werden:

| Feld | Beschreibung |
|------|--------------|
| **Name** | Produktname (Pflichtfeld) |
| **Kategorie** | Zuordnung zu einer Kategorie |
| **Preis** | Preis in Euro (optional) |
| **Shop-Link** | URL zum Shop (optional) |
| **Notizen** | Zusätzliche Informationen (optional) |

## Kategorien

Standardmäßig sind folgende Kategorien verfügbar:

1. Elektrik
2. Wasser
3. Küche
4. Schlafen
5. Stauraum

## Datenformat

Die Daten werden im JSON-Format gespeichert:

```json
{
  "categories": ["Elektrik", "Wasser", "Küche", "Schlafen", "Stauraum"],
  "products": [
    {
      "id": "uuid",
      "name": "Produktname",
      "category": "Elektrik",
      "completed": false,
      "price": 99.99,
      "shopLink": "https://...",
      "notes": "Notizen",
      "customFields": {}
    }
  ]
}
```

## Tastatur-Shortcuts (Übersicht)

```
Navigation:
  j/k       - Runter/Hoch
  ←/→       - Kategorie wechseln
  1-5       - Direkte Kategorie-Auswahl

Aktionen:
  Space     - Toggle erledigt
  n         - Neues Produkt
  e         - Bearbeiten
  d         - Löschen
  /         - Suche

Allgemein:
  q         - Quit
  Tab/Enter - Formular weiter
  Esc       - Abbrechen
```

## Lizenz

MIT License