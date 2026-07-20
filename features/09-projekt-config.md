# Feature: Projekt-Config

## Status
✅ Implementiert

## Beschreibung
Jedes Projekt erhält eine `config.json` Datei, die Projekteinstellungen speichert. Beim Start eines Projekts wird automatisch das zuletzt verwendete Produkt geladen.

## Funktionalität

### Config-Datei erstellen
- Beim Erstellen eines neuen Projekts wird automatisch eine `config.json` mit erstellt
- Config wird im gleichen Verzeichnis wie die Projekt-JSON gespeichert
- Standard-Config mit leerem `lastProductId`

### Zuletzt verwendetes Produkt speichern
- Beim Auswählen/Wechseln eines Produkts wird die Product-ID in der Config gespeichert
- Config wird automatisch bei Änderung aktualisiert

### Produkt beim Start laden
- Beim Laden eines Projekts wird die Config gelesen
- Falls ein `lastProductId` gespeichert ist, wird dieser automatisch ausgewählt
- Falls das Produkt nicht mehr existiert, wird kein Produkt vorausgewählt

## Dateien
- `internal/storage/project.go` – Config-Les/Schreib-Logik erweitern
- `internal/model/config.go` – Neues Config-Modell
- `internal/ui/app.go` – Auto-Load-Logik beim Projektstart

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| - | Keine neuen Shortcuts |

## Datenmodell
```json
{
  "lastProductId": "uuid-string"
}
```

## Abhängigkeiten
- Feature 08 (Projekte) muss implementiert sein
- Feature 06 (JSON-Persistenz) muss implementiert sein