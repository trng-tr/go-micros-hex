package dtos

import "time"

type OrderLineRequest struct {
	ProductID  int64 `json:"product_id" binding:"required"`
	LocationID int64 `json:"location_id" binding:"required"`
	Quantity   int64 `json:"quantity" binding:"required"`
}
type OrderRequest struct {
	CustomerID int64              `json:"customer_id" binding:"required"`
	OrderLines []OrderLineRequest `json:"order_lines"`
}

type OrderResponse struct {
	ID                  int64                 `json:"id"`
	CustomerResponse    LightCustomerResponse `json:"customer"`
	Status              string                `json:"order_status"`
	OrderLinesResponses []OrderLineResponse   `json:"lines"`
	CreatedAt           string                `json:"created_at"`
}

type OrderLineResponse struct {
	ID               int64            `json:"id"`
	ProductResponse  ProductResponse  `json:"product"`
	LocationResponse LocationResponse `json:"location"`
	Quantity         int64            `json:"quantity"`
}

type Response struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

func NewResponse(status, message string) *Response {
	return &Response{
		Status:    status,
		Message:   message,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
}

type AjustQuantityRequest struct {
	Quantity int64 `json:"quantity" binding:"required"`
}
