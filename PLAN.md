# Plan: TUI-Überarbeitung

## Zusammenfassung

Die TUI soll in 4 Punkten überarbeitet werden:
1. Checkbox-Rendering von Unicode-Symbolen zu ASCII `[ ]` / `[x]`
2. Entfernung der Leerzeilen zwischen Kategorien
3. Entfernung der Detailansicht für ausgewählte Produkte
4. Hinzufügen einer Leerzeile zwischen Header und Content

---

## Phase 1: Feature-Dateien anpassen

### Datei: `features/01-tui.md`

#### Änderung 1: UI-Beispiel aktualisieren
**Abschnitt:** UI-Beispiel (Zeile 81-99)

**Alt:**
```
☐ Reifen
☑ Batteriewechsel
```

**Neu:**
```
[ ] Reifen
[x] Batteriewechsel
```

#### Änderung 2: Tastatur-Shortcuts prüfen
**Abschnitt:** Listenansicht (Zeile 51-61)

Keine Änderungen erforderlich, Shortcuts bleiben gleich.

#### Änderung 3: Beschreibung ergänzen
**Abschnitt:** Beschreibung (Zeile 6-7)

**Neuer Text:**
```
Zentrales TUI-Layout mit Header, scrollbarem Content-Bereich und kontext-sensitivem Footer.
Produktliste wird kategorisiert ohne Detailansicht angezeigt.
Jede View liefert ihre eigene Tastaturnavigation über eine `GetShortcuts()`-Methode.
```

---

## Phase 2: Code-Änderungen

### Datei: `internal/ui/project_view.go`

#### Änderung 1: Checkbox-Rendering (Zeile 114-119)

**Ort:** Methode `renderCategory()`

**Alt:**
```go
checkbox := "☐"
if product.Completed {
    checkbox = "☑"
}

line := fmt.Sprintf("%s[%s] %d. %s", cursor, checkbox, product.Index, product.Name)
```

**Neu:**
```go
checkbox := "[ ]"
if product.Completed {
    checkbox := "[x]"
}

line := fmt.Sprintf("%s%s %d. %s", cursor, checkbox, product.Index, product.Name)
```

---

#### Änderung 2: Leerzeile zwischen Kategorien entfernen (Zeile 88-91)

**Ort:** Methode `renderContent()`

**Alt:**
```go
for _, category := range p.project.Categories {
    products, exists := categoryProducts[category]
    if !exists || len(products) == 0 {
        continue
    }

    sb.WriteString(p.renderCategory(category, products))
    sb.WriteString("\n")
}
```

**Neu:**
```go
for _, category := range p.project.Categories {
    products, exists := categoryProducts[category]
    if !exists || len(products) == 0 {
        continue
    }

    sb.WriteString(p.renderCategory(category, products))
}
```

---

#### Änderung 3: Detailansicht entfernen (Zeile 93-96)

**Ort:** Methode `renderContent()`

**Alt:**
```go
if len(p.project.Products) > 0 {
    sb.WriteString(p.renderDetailSection())
}

return sb.String()
```

**Neu:**
```go
return sb.String()
```

**Zusätzlich:** Methode `renderDetailSection()` (Zeile 143-178) kann entfernt werden, da sie nicht mehr verwendet wird.

**Zusätzlich:** Methode `statusText()` (Zeile 180-185) kann entfernt werden, da sie nur von `renderDetailSection()` verwendet wurde.

---

#### Änderung 4: Leerzeile zwischen Header und Content (Zeile 71-73)

**Ort:** Methode `View()`

**Alt:**
```go
func (p *ProjectView) View() string {
    return p.viewport.View()
}
```

**Neu:**
```go
func (p *ProjectView) View() string {
    return "\n" + p.viewport.View()
}
```

---

## Phase 3: Testen

### Manuelle Tests

1. **Checkbox-Rendering**
   - [ ] Nicht gecheckte Items zeigen `[ ]`
   - [ ] Gecheckte Items zeigen `[x]`
   - [ ] Space-Taste toggelt Checkbox korrekt

2. **Kategorie-Abstände**
   - [ ] Keine Leerzeilen zwischen Kategorien
   - [ ] Kategorien folgen direkt aufeinander

3. **Detailansicht**
   - [ ] Keine Detailansicht unter der Produktliste
   - [ ] Mehr Platz für Produktliste verfügbar

4. **Header-Abstand**
   - [ ] Eine Leerzeile zwischen Header und Content
   - [ ] Layout insgesamt korrekt

### Build-Test
```bash
go build -o van-planner .
./van-planner
```

---

## Dateien-Übersicht

| Datei | Änderungen |
|-------|------------|
| `features/01-tui.md` | UI-Beispiel, Beschreibung |
| `internal/ui/project_view.go` | `renderCategory()`, `renderContent()`, `View()`, Methoden entfernen |

---

## Abhängigkeiten

- Keine neuen Dependencies
- Bestehende Libraries: bubbletea, lipgloss, bubbles

---

## Risiken

- **Niedrig**: Alle Änderungen sind kosmetisch
- Keine Änderungen am Datenmodell
- Keine Änderungen an der Geschäftslogik

---

## Geschätzter Aufwand

- Feature-Dateien: 5 Minuten
- Code-Änderungen: 15 Minuten
- Testing: 10 Minuten
- **Gesamt: ~30 Minuten**