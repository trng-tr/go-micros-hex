package in

import (
	"context"

	"github.com/trng-tr/product-microservice/internal/domain"
)

// InProductService de gestion des produits
type InProductService interface {
	SaveProduct(ctx context.Context, prd domain.Product) (domain.Product, error)
	GetProductByID(ctx context.Context, id int64) (domain.Product, error)
	GetAllProducts(ctx context.Context) ([]domain.Product, error)
	PatchProduct(ctx context.Context, id int64, patchProduct domain.PatchProduct) (domain.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
	GetProductBySku(ctx context.Context, sku string) (domain.Product, error)
}

// InStockService de gestion du stock pour un produit
type InStockService interface {
	CreateStock(ctx context.Context, stk domain.Stock) (domain.Stock, error)
	GetStockByID(ctx context.Context, id int64) (domain.Stock, error)
	GetAllStocks(ctx context.Context) ([]domain.Stock, error)
	SetStockQuantity(ctx context.Context, productID int64, newQuantity int64) (domain.Stock, error)   //replace quantity
	IncreaseStockQuantity(ctx context.Context, productID int64, quantity int64) (domain.Stock, error) // encrease
	DecreaseStockQuantity(ctx context.Context, productID int64, quantity int64) (domain.Stock, error) //decrease
	GetStockByProductID(ctx context.Context, productID int64) (domain.Stock, error)
}

// InLocationService chaque stock est localisée dans une ville
type InLocationService interface {
	CreateLocation(ctx context.Context, location domain.Location) (domain.Location, error)
	GetLocationByID(ctx context.Context, id int64) (domain.Location, error)
	GetAllLocation(ctx context.Context) ([]domain.Location, error)
}
