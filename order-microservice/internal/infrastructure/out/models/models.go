package models

import "time"

type OrderModel struct {
	ID         int64
	CustomerID int64 //remote api: 8081
	CreatedAt  time.Time
	Status     string
	Lines      []OrderLineModel
}

type OrderLineModel struct {
	ID        int64
	OrderID   int64
	ProductID int64 //remote api: 8082
	Quantity  int64
}
