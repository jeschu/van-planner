package model

import "github.com/google/uuid"

type Product struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Category     string                 `json:"category"`
	Completed    bool                   `json:"completed"`
	Price        float64                `json:"price"`
	ShopLink     string                 `json:"shopLink"`
	Notes        string                 `json:"notes"`
	CustomFields map[string]interface{} `json:"customFields"`
}

func NewProduct(name, category string) Product {
	return Product{
		ID:           uuid.New().String(),
		Name:         name,
		Category:     category,
		Completed:    false,
		Price:        0.0,
		ShopLink:     "",
		Notes:        "",
		CustomFields: make(map[string]interface{}),
	}
}

type Data struct {
	Categories []string  `json:"categories"`
	Products   []Product `json:"products"`
}

func NewData() Data {
	return Data{
		Categories: []string{"Elektrik", "Wasser", "Küche", "Schlafen", "Stauraum"},
		Products:   []Product{},
	}
}

type ProjectConfig struct {
	LastProductID string `json:"lastProductId,omitempty"`
}

func NewProjectConfig() ProjectConfig {
	return ProjectConfig{
		LastProductID: "",
	}
}
