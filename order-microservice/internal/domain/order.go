package domain

import "time"

type Order struct {
	ID         int64
	CustomerID int64 //remote api: 8081
	CreatedAt  time.Time
	Status     OrderStatus
	Lines      []OrderLine
}

type OrderLine struct {
	ID        int64
	OrderID   int64
	ProductID int64 //remote api: 8082
	Quantity  int64
}

type OrderStatus string

const (
	Created   OrderStatus = "CREATED"
	Confirmed OrderStatus = "CONFIRMED"
	Payed     OrderStatus = "PAID"
)
