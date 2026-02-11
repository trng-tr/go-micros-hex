package dtos

type LocationResponse struct {
	ID          int64   `json:"id"`
	Ville       string  `json:"ville"`
	Description *string `json:"description,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}
type LightLocationResponse struct {
	ID          int64   `json:"id"`
	Ville       string  `json:"ville"`
	Description *string `json:"description,omitempty"`
}

type StockResponse struct {
	ID                    int64                 `json:"id"`
	Name                  string                `json:"stock_name"`
	LocationID            int64                 `json:"location_id"`
	LightLocationResponse LightLocationResponse `json:"location"`
	ProductID             int64                 `json:"product_id"`
	ProductResponse       ProductResponse       `json:"product"`
	Quantity              int64                 `json:"stock_quantity"`
}
