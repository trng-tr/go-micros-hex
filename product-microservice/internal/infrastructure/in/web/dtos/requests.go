package dtos

type ProductRequest struct {
	Category     string       `json:"category" binding:"required"`
	ProductName  string       `json:"product_name" binding:"required"`
	Description  string       `json:"description" binding:"required"`
	PriceRequest PriceRequest `json:"price" binding:"required"`
}

type PriceRequest struct {
	UnitPrice int64  `json:"unit_price" binding:"required"`
	Currency  string `json:"currency" binding:"required"`
}

type ProductPatchRequest struct {
	ProductName *string `json:"product_name,omitempty" binding:"omitempty,min=2"`
	Description *string `json:"description,omitempty" binding:"omitempty,min=2"`
	UnitPrice   *int64  `json:"unit_price,omitempty" binding:"omitempty,gt=0"`
}

type Locationrequest struct {
	Ville       string  `json:"ville" binding:"required"`
	Description *string `json:"description,omitempty"`
}
type StockRequest struct {
	Name       string `json:"stock_name"`
	ProductID  int64  `json:"product_id"`
	LocationID int64  `json:"location_id"`
	Quantity   int64  `json:"quantity"`
}

// StockQuantityRequest struct to set,increase or decrease quantity of a stock
type StockQuantityRequest struct {
	Quantity int64 `json:"quantity"`
}
