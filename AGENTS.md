# AI-Assistent Anweisungen

## Sprache
- Antworte immer auf Deutsch, es sei denn der User fragt explizit nach einer anderen Sprache.
- Code-Kommentare und Commit-Messages entsprechend dem Projektkontext wählen.

## Kommunikation
- Sei prägnant und direkt.
- Vermeide unnötige Erklärungen und Zusammenfassungen.
- Keine Emojis, es sei denn explizit gewünscht.
- Wenn du dir unsicher bist, haluziniere nicht sondern frage nach.
- Der User heißt Jens. Sprich ihn mit seinem Namen an.
- Wenn Rückfragen gestellt werden, nutze das Frage-Tool.

## Code
- Folge bestehenden Konventionen des Projekts.
- Bevorzuge Editing bestehender Dateien über das Erstellen neuer Dateien.
- Schreibe keine Kommentare im Code, es sei denn explizit gewünscht.
- Verwende niemals Bibliotheken, ohne zuerst zu prüfen ob sie bereits im Projekt verwendet werden.

## Sicherheit
- Committe niemals Secrets, Keys oder Credentials.
- Logge niemals sensible Daten.

## Tools
- Verwende parallele Tool-Calls wo möglich.
- Runne Lint/Typecheck-Befehle nach Code-Änderungen, falls verfügbar.
- Vermeide head/tail/grep - verwende die dedizierten Tools stattdessen.

## Feature-Driven Development

### Feature-Dateien (`features/`)
Features werden in nummerierten Markdown-Dateien im `features/` Ordner spezifiziert.

#### Struktur einer Feature-Datei
```markdown
# Feature: [Feature-Name]

## Status
🔲 Geplant / ⬇️ In Arbeit / ✅ Implementiert

## Beschreibung
Kurze Beschreibung des Features.

## Funktionalität
Detaillierte Beschreibung der Features und Funktionen.

## Dateien
- `pfad/zur/datei.go` – Beschreibung

## Tastatur-Shortcuts
| Taste | Aktion |
|-------|--------|
| `x` | Aktion |

## Datenmodell
```json
{...}
```
```

### Workflow

#### Neues Feature implementieren
1. Neue Feature-Datei mit nächster Nummer erstellen (z.B. `07-neues-feature.md`)
2. Feature nach Vorlage (`00-vorlage.md`) spezifizieren
3. Datei committen
4. Implementierung basierend auf Feature-Datei durchführen

#### Feature anpassen
1. Existierende Feature-Datei editieren
2. Änderungen committen
3. Code-Anpassungen basierend auf aktualisierter Datei vornehmen

#### Regeln
- Feature-Dateien sind die **Single Source of Truth** für Anforderungen
- Vor Implementierung: Feature-Datei lesen und verstehen
- Nach Implementierung: Feature-Datei auf `✅ Implementiert` setzen
- Bei Unklarheiten in Feature-Datei: Nachfragen statt raten

### Branch-Strategie
- Feature-Branches: `feature/<name>` (z.B. `feature/phase1-grundgeruest`)
- Pro Feature-Datei kann ein eigener Branch erstellt werden
- Nach Abschluss: Merge in Hauptbranch
