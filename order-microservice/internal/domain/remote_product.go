package domain

type Product struct {
	ID          int64
	Sku         string
	Category    Category
	ProductName string
	Description string
	Price       Price
	IsActive    bool
}

type Category string

const (
	Book     Category = "BOOK"
	Clothing Category = "CLTH"
	Shoes    Category = "SHOE"
)

type Price struct {
	UnitPrice int64
	Currency  Currency
}
type Currency string

const (
	Dollar Currency = "USD"
	Euro   Currency = "EUR"
)

type Stock struct {
	ID        int64
	Name      string
	ProductID int64
	Quantity  int64
}
