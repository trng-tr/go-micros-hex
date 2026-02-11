package mappers

import (
	"database/sql"
	"time"

	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/out/models"
)

func ToProductModel(prod domain.Product) models.ProductModel {
	var updatedAt sql.NullTime
	if prod.UpdatedAt != nil {
		updatedAt = sql.NullTime{
			Time:  *prod.UpdatedAt,
			Valid: true,
		}
	}
	return models.ProductModel{
		ID:          prod.ID,
		Sku:         prod.Sku,
		Category:    string(prod.Category),
		ProductName: prod.ProductName,
		Description: prod.Description,
		UnitPrice:   prod.Price.UnitPrice,
		Currency:    string(prod.Price.Currency),
		CreatedAt:   prod.CreatedAt,
		UpdatedAt:   updatedAt,
		IsActive:    prod.IsActive,
	}
}

func ToBusinessProduct(model models.ProductModel) domain.Product {
	var updatedAt *time.Time
	if model.UpdatedAt.Valid {
		updatedAt = &model.UpdatedAt.Time
	} else {
		updatedAt = nil
	}
	return domain.Product{
		ID:          model.ID,
		Sku:         model.Sku,
		Category:    domain.Category(model.Category),
		ProductName: model.ProductName,
		Description: model.Description,
		Price: domain.Price{
			UnitPrice: model.UnitPrice,
			Currency:  domain.Currency(model.Currency),
		},
		CreatedAt: model.CreatedAt,
		UpdatedAt: updatedAt,
		IsActive:  model.IsActive,
	}
}
