# Feature: JSON-Persistenz

## Status
✅ Implementiert

## Beschreibung
Lokale Speicherung aller Daten im JSON-Format.

## Funktionalität

### Speicherort
- Standard: `data/products.json`
- Automatische Verzeichnis-Erstellung

### Save-On-Change
- Automatische Persistenz nach jeder Änderung
- Create, Update, Delete, Toggle

### Load-On-Start
- Laden beim Anwendungsstart
- Leere Daten bei nicht-existierender Datei

### JSON-Struktur
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

## Dateien
- `internal/storage/json.go` – Storage-Implementierung
- `data/products.json` – Datenspeicher

## Schnittstelle
```go
type JSONStorage struct {
    filepath string
}

func (s *JSONStorage) Load() (model.Data, error)
func (s *JSONStorage) Save(data model.Data) error
```

## Fehlerbehandlung
- File not found → Leere Daten
- Invalid JSON → Fehlermeldung
- Write error → Toast-Nachricht in UI