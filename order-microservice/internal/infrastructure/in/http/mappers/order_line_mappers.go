package mappers

import (
	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/dtos"
)

func ToBusinessOrderLine(request dtos.OrderLineRequest) domain.OrderLine {
	return domain.OrderLine{
		ProductID:  request.ProductID,
		LocationID: request.LocationID,
		Quantity:   request.Quantity,
	}
}

func ToOrderLineResponse(orderLine domain.OrderLine, product domain.Product, location domain.Location) dtos.OrderLineResponse {
	return dtos.OrderLineResponse{
		ID:               orderLine.ID,
		ProductResponse:  ToProductResponse(product),
		LocationResponse: ToLocationResponse(location),
		Quantity:         orderLine.Quantity,
	}
}
