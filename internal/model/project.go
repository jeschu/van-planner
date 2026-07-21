package model

type Project struct {
	Categories []string  `json:"categories"`
	Products   []Product `json:"products"`
}

type Product struct {
	Index         int            `json:"index"`
	Name          string         `json:"name"`
	Count         *int           `json:"count"`
	EstimatedCost float64        `json:"estimatedCost"`
	ActualCost    float64        `json:"actualCost"`
	ShippingCost  float64        `json:"shippingCost"`
	ShopLink      string         `json:"shopLink"`
	Notes         string         `json:"notes"`
	Completed     bool           `json:"completed"`
	Category      string         `json:"category"`
	CustomFields  map[string]any `json:"customFields"`
}
