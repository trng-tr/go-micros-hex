package mappers

import (
	"time"

	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/dtos"
)

func ToStockResponse(bsStock domain.Stock, bsProduct domain.Product, bsLocation domain.Location) dtos.StockResponse {
	return dtos.StockResponse{
		ID:         bsStock.ID,
		Name:       bsStock.Name,
		Quantity:   bsStock.Quantity,
		LocationID: bsLocation.ID,
		LightLocationResponse: dtos.LightLocationResponse{
			ID:          bsLocation.ID,
			Ville:       bsLocation.Ville,
			Description: bsLocation.Description,
		},
		ProductID:       bsProduct.ID,
		ProductResponse: ToProductResponse(bsProduct),
	}
}

func ToBusinessStock(request dtos.StockRequest) domain.Stock {
	return domain.Stock{
		Name:       request.Name,
		ProductID:  request.ProductID,
		LocationID: request.LocationID,
		Quantity:   request.Quantity,
	}
}

func ToLocationResponse(bsLocation domain.Location) dtos.LocationResponse {
	var updatedAt *string
	if bsLocation.UpdatedAt != nil {
		var s = bsLocation.UpdatedAt.Format(time.RFC3339)
		updatedAt = &s
	}
	return dtos.LocationResponse{
		ID:          bsLocation.ID,
		Ville:       bsLocation.Ville,
		Description: bsLocation.Description,
		CreatedAt:   bsLocation.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   updatedAt,
	}
}

func ToBusinessLocation(dto dtos.Locationrequest) domain.Location {
	return domain.Location{
		Ville:       dto.Ville,
		Description: dto.Description,
	}
}
