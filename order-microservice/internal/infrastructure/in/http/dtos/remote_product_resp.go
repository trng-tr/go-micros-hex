package dtos

type ProductResponse struct {
	ID            int64         `jsons:"id"`
	Sku           string        `json:"sku"`
	Category      string        `json:"category"`
	ProductName   string        `json:"product_name"`
	Description   string        `json:"description"`
	PriceResponse PriceResponse `json:"price"`
	IsActive      bool          `json:"is_active"`
}

type PriceResponse struct {
	UnitPrice int64  `json:"unit_price"`
	Currency  string `json:"currency"`
}
type LocationResponse struct {
	ID          int64   `json:"id"`
	Ville       string  `json:"ville"`
	Description *string `json:"description,omitempty"`
}
type StockResponse struct {
	ID               int64            `json:"id"`
	Name             string           `json:"name"`
	ProductID        int64            `json:"product_id"`
	LocationID       int64            `json:"location_id"`
	Quantity         int64            `json:"stock_quantity"`
	UpdatedAt        string           `json:"updated_at"`
	LocationResponse LocationResponse `json:"location"`
	ProductResponse  ProductResponse  `json:"product"`
}
