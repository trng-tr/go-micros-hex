package models

import (
	"database/sql"
	"time"
)

type ProductModel struct {
	ID          int64
	Sku         string
	Category    string
	ProductName string
	Description string
	UnitPrice   int64
	Currency    string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime // it not mandaroty
	IsActive    bool
}

type StockModel struct {
	ID         int64
	Name       string
	ProductID  int64
	LocationID int64 //ville of stock
	Quantity   int64
	UpdatedAt  time.Time
}

type Location struct {
	ID          int64
	Ville       string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime // it not mandaroty
}
