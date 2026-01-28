package domain

import (
	"fmt"
	"strings"
	"time"
)

type Product struct {
	ID          int64
	Sku         string
	Category    Category
	ProductName string
	Description string
	Price       Price
	CreatedAt   time.Time
	UpdatedAt   *time.Time // it not mandaroty
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
type PatchProduct struct {
	ProductName *string
	Description *string
	UnitPrice   *int64
}

type Currency string

const (
	Dollar Currency = "USD"
	Euro   Currency = "EUR"
)

func (p *Product) GenerateSku(prodCategory Category, prodName string, uuid string) {
	var cuttedName string
	if len(strings.TrimSpace(prodName)) > 4 {
		cuttedName = prodName[:5]
	} else {
		cuttedName = prodName
	}
	p.Sku = fmt.Sprintf("%s-%s-%s", prodCategory, cuttedName, uuid)
}

// GenerateCreatedAt util to set created date to product
func (p *Product) GenerateCreatedAt() {
	p.CreatedAt = time.Now()
}

// GenerateUpdateAt util to set updated date to product
func (p *Product) GenerateUpdateAt() {
	var now = time.Now()
	p.UpdatedAt = &now
}

// ApplyPatchMapper mapper for pacth method
func (p *Product) ApplyPatchMapper(patch PatchProduct) {
	if patch.ProductName != nil {
		p.ProductName = *patch.ProductName
	}
	if patch.Description != nil {
		p.Description = *patch.Description
	}
	if patch.UnitPrice != nil {
		p.Price.UnitPrice = *patch.UnitPrice
	}

	p.GenerateUpdateAt() //update date value
}
