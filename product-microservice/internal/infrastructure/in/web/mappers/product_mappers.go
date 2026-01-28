package mappers

import (
	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/dtos"
)

const dateFormate string = "2006-01-02 15:04:05"

// ToBusinessProduct mapper for create request
func ToBusinessProduct(request dtos.ProductRequest) domain.Product {
	return domain.Product{
		Category:    domain.Category(request.Category),
		ProductName: request.ProductName,
		Description: request.Description,
		Price: domain.Price{
			UnitPrice: request.PriceRequest.UnitPrice,
			Currency:  domain.Currency(request.PriceRequest.Currency),
		},
	}
}

// ToProductResponse mapper for user response
func mapCurrencyToSymbol(bsProduct domain.Product) string {
	switch bsProduct.Price.Currency {
	case domain.Dollar:
		return "$"
	case domain.Euro:
		return "€"
	default:
		return string(bsProduct.Price.Currency)
	}
}
func ToProductResponse(bsProduct domain.Product) dtos.ProductResponse {
	var priceRensponse dtos.PriceResponse = dtos.PriceResponse{
		UnitPrice: bsProduct.Price.UnitPrice,
		Currency:  mapCurrencyToSymbol(bsProduct),
	}

	var updatedAt *string
	if bsProduct.UpdatedAt != nil {
		s := bsProduct.UpdatedAt.Format(dateFormate)
		updatedAt = &s
	}
	return dtos.ProductResponse{
		ID:            bsProduct.ID,
		Sku:           bsProduct.Sku,
		Category:      string(bsProduct.Category),
		ProductName:   bsProduct.ProductName,
		Description:   bsProduct.Description,
		PriceResponse: priceRensponse,
		CreatedAt:     bsProduct.CreatedAt.Format(dateFormate),
		UpdatedAt:     updatedAt,
		IsActive:      bsProduct.IsActive,
	}
}

func ToDomainPatchRequest(patchRequest dtos.ProductPatchRequest) domain.PatchProduct {
	return domain.PatchProduct{
		ProductName: patchRequest.ProductName,
		Description: patchRequest.Description,
		UnitPrice:   patchRequest.UnitPrice,
	}
}
