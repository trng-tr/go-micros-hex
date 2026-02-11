package mappers

import (
	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/dtos"
)

func ToCustomerResponse(customer domain.Customer) dtos.LightCustomerResponse {
	return dtos.LightCustomerResponse{
		ID:          customer.ID,
		Firstname:   customer.Firstname,
		Lastname:    customer.Lastname,
		Genda:       string(customer.Genda),
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Status:      string(customer.Status),
	}
}

func ToProductResponse(bsProduct domain.Product) dtos.ProductResponse {
	return dtos.ProductResponse{
		ID:          bsProduct.ID,
		Sku:         bsProduct.Sku,
		Category:    string(bsProduct.Category),
		ProductName: bsProduct.ProductName,
		Description: bsProduct.Description,
		PriceResponse: dtos.PriceResponse{
			UnitPrice: bsProduct.Price.UnitPrice,
			Currency:  string(bsProduct.Price.Currency),
		},
		IsActive: bsProduct.IsActive,
	}
}

func ToLocationResponse(bsLocation domain.Location) dtos.LocationResponse {
	return dtos.LocationResponse{
		ID:          bsLocation.ID,
		Ville:       bsLocation.Ville,
		Description: bsLocation.Description,
	}
}
