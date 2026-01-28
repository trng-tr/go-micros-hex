package dtos

type LocationResponse struct {
	ID          int64   `json:"id"`
	Ville       string  `json:"ville"`
	Description *string `json:"description,omitempty"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}
type LightLocationResponse struct {
	Ville       string  `json:"ville"`
	Description *string `json:"description,omitempty"`
}

type StockResponse struct {
	ID                    int64                 `json:"id"`
	Name                  string                `json:"stock_name"`
	LightLocationResponse LightLocationResponse `json:"location"`
	ProductResponse       ProductResponse       `json:"product"`
	Quantity              int64                 `json:"stock_quantity"`
}
