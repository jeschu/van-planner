package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/jens/van-planner/internal/model"
)

type JSONStorage struct {
	filepath string
}

func NewJSONStorage(path string) *JSONStorage {
	return &JSONStorage{
		filepath: path,
	}
}

type legacyProduct struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Category     string                 `json:"category"`
	Completed    bool                   `json:"completed"`
	Price        float64                `json:"price"`
	ShopLink     string                 `json:"shopLink"`
	Notes        string                 `json:"notes"`
	CustomFields map[string]interface{} `json:"customFields"`
}

type legacyData struct {
	Categories []string        `json:"categories"`
	Products   []legacyProduct `json:"products"`
}

func (s *JSONStorage) Load() (model.Data, error) {
	data := model.NewData()

	content, err := os.ReadFile(s.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return data, nil
		}
		return data, err
	}

	var legacy legacyData
	if err := json.Unmarshal(content, &legacy); err == nil {
		data.Categories = legacy.Categories
		for _, p := range legacy.Products {
			product := model.Product{
				ID:            p.ID,
				Name:          p.Name,
				Category:      p.Category,
				Completed:     p.Completed,
				EstimatedCost: p.Price,
				ActualCost:    0,
				ShopLink:      p.ShopLink,
				Notes:         p.Notes,
				CustomFields:  p.CustomFields,
			}
			data.Products = append(data.Products, product)
		}
		if len(data.Products) > 0 {
			return data, nil
		}
	}

	if err := json.Unmarshal(content, &data); err != nil {
		return data, err
	}

	return data, nil
}

func (s *JSONStorage) Save(data model.Data) error {
	dir := filepath.Dir(s.filepath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filepath, content, 0644)
}
