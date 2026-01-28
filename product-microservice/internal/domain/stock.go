package domain

import "time"

type Stock struct {
	ID         int64
	Name       string
	ProductID  int64
	LocationID int64 //ville of stock
	Quantity   int64
	UpdatedAt  time.Time
}
