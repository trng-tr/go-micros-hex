package dtos

import "time"

type OrderLineRequest struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int64 `json:"quantity" binding:"required"`
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
	ID              int64
	ProductResponse ProductResponse `json:"product"`
	Quantity        int64
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

type AjustStockQuantityRequest struct {
	Quantity int64 `json:"quantity" binding:"required"`
}
