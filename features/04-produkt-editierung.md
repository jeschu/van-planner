# Feature: Produkt-Editierung

## Status
✅ Implementiert

## Beschreibung
Auf der Produktseite kann das gewählte Produkt mit `e` editiert werden. Ein eigenes Formular ermöglicht die Bearbeitung aller Produktfelder. Mit `CTRL-S` wird gespeichert und man kehrt zur Hauptseite zurück (aktuelles Produkt bleibt gewählt, Scroll-Position erhalten). Mit `CTRL-D` kann ein Produkt im Edit-Modus gelöscht werden (mit Bestätigungsdialog). Mit `n` wird auf der Hauptseite ein neues Produkt unter dem aktuellen Produkt hinzugefügt.

## Funktionalität

### Edit-Modus starten
- Taste `e` auf der Hauptseite startet den Edit-Modus für das aktuell ausgewählte Produkt
- Ein separates Formular wird angezeigt mit allen editierbaren Feldern

### Produkt-Formular
- Alle Produktfelder können bearbeitet werden:
  - Name (Text)
  - Kategorie (Dropdown/Text)
  - Geschätzte Kosten (Zahl)
  - Tatsächliche Kosten (Zahl)
  - Versandkosten (Zahl)
  - Anzahl (Zahl, optional)
  - Shop-Link (Text)
  - Notizen (Text, mehrzeilig)
- Navigation zwischen Feldern mit `Tab` / `Shift-Tab` oder `↑` / `↓`
- Eingabe von Text/Zahlen direkt im Formular möglich

### Speichern (CTRL-S)
- `CTRL-S` speichert das Produkt
- Projekt wird persistiert (JSON-Datei)
- Rückkehr zur Hauptseite
- Das gleiche Produkt bleibt ausgewählt
- Scroll-Position bleibt erhalten

### Löschen (CTRL-D)
- `CTRL-D` im Edit-Modus zeigt Bestätigungsdialog
- Dialog fragt "Produkt wirklich löschen?" mit Ja/Nein-Optionen
- Bei Bestätigung: Produkt wird entfernt und Projekt gespeichert
- Bei Abbruch: Dialog schließt sich ohne Änderung

### Neues Produkt (n)
- Taste `n` auf der Hauptseite fügt neues Produkt unter dem aktuellen Produkt ein
- Neues Produkt erhält nächstfreien Index
- Kategorie des aktuellen Produkts wird übernommen
- Edit-Modus öffnet sich automatisch für das neue Produkt

## Dateien
- `internal/ui/product_edit.go` – Formular zur Produktbearbeitung
- `internal/ui/delete_dialog.go` – Bestätigungsdialog zum Löschen
- `internal/ui/project_view.go` – Erweiterung um 'n' für neues Produkt
- `internal/ui/app.go` – Integration des Edit-States
- `internal/storage/project.go` – Speichern der Änderungen

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `e` | Edit-Modus für ausgewähltes Produkt starten |
| `CTRL-S` | Produkt speichern und zur Hauptseite zurückkehren |
| `CTRL-D` | Produkt löschen (mit Bestätigungsdialog) |
| `n` | Neues Produkt unter dem aktuellen einfügen |
| `Tab` / `↑` | Nächstes Feld im Formular |
| `Shift-Tab` / `↓` | Vorheriges Feld im Formular |
| `Esc` | Bearbeiten abbrechen (nur im Edit-Modus) |

## Datenmodell
Keine Änderungen am Datenmodell erforderlich.

## UI-Beispiel (Edit-Modus)
```
┌─────────────────────────────────────────────────────────────┐
│  Van Planner - Projekt: …                      [EDIT MODE]  │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Produkt bearbeiten                                         │
│                                                             │
│  Name:        [Ducato                    ]                  │
│  Kategorie:   [Fahrzeug                  ]                  │
│  Geschätzte Kosten:   [0.00              ] €                │
│  Tatsächliche Kosten: [0.00              ] €                │
│  Versandkosten:       [0.00              ] €                │
│  Anzahl:              [                    ]                  │
│  Shop-Link:           [https://          ]                  │
│                                                             │
│  Notizen:                                                   │
│  ┌─────────────────────────────────────────────────────┐   │
│  │                                                     │   │
│  │                                                     │   │
│  └─────────────────────────────────────────────────────┘   │
│                                                             │
│  [CTRL-S] Speichern  [CTRL-D] Löschen  [Esc] Abbrechen     │
├─────────────────────────────────────────────────────────────┤
│  [↑]↑ [↓]↓ Navigieren [Tab] Feld wechseln                  │
└─────────────────────────────────────────────────────────────┘
```

## UI-Beispiel (Lösch-Dialog)
```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  ┌───────────────────────────────────────────────────────┐ │
│  │  Produkt wirklich löschen?                            │ │
│  │                                                       │ │
│  │  [Ja]  [Nein]                                         │ │
│  │                                                       │ │
│  └───────────────────────────────────────────────────────┘ │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

## Abhängigkeiten
- Feature 01 (TUI) muss implementiert sein
- Feature 02 (Projekt) muss implementiert sein
- Bubbles Text Input (bereits in go.mod)
- Bubbles Text Area (bereits in go.mod)

## Offene Fragen
- Soll die Kategorie aus einer festen Liste ausgewählt werden oder frei editierbar sein?
- Sollen Zahlenfelder mit Validierung versehen werden?