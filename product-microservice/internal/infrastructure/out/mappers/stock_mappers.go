package mappers

import (
	"database/sql"
	"time"

	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/out/models"
)

// ToStockModel mapper to object for db
func ToStockModel(bsStock domain.Stock) models.StockModel {

	return models.StockModel{
		ID:         bsStock.ID,
		Name:       bsStock.Name,
		ProductID:  bsStock.ProductID,
		LocationID: bsStock.LocationID,
		Quantity:   bsStock.Quantity,
		UpdatedAt:  bsStock.UpdatedAt,
	}
}

func ToBusinessStock(model models.StockModel) domain.Stock {
	return domain.Stock{
		ID:         model.ID,
		Name:       model.Name,
		ProductID:  model.ProductID,
		LocationID: model.LocationID,
		Quantity:   model.Quantity,
		UpdatedAt:  model.UpdatedAt,
	}
}

func ToLocationModel(location domain.Location) models.Location {
	var updatedAt sql.NullTime
	if location.UpdatedAt != nil {
		updatedAt = sql.NullTime{
			Time:  *location.UpdatedAt,
			Valid: true,
		}
	}
	return models.Location{
		ID:          location.ID,
		Ville:       location.Ville,
		Description: location.Description,
		CreatedAt:   location.CreatedAt,
		UpdatedAt:   updatedAt,
	}
}

func ToBusinessLocation(model models.Location) domain.Location {
	var updatedAt *time.Time
	if model.UpdatedAt.Valid {
		updatedAt = &model.UpdatedAt.Time
	}
	return domain.Location{
		ID:          model.ID,
		Ville:       model.Ville,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   updatedAt,
	}
}
