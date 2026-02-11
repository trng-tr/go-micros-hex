package mappers

import (
	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/out/models"
)

// order mappers
func ToOrderModel(order domain.Order) models.OrderModel {
	var lines []models.OrderLineModel
	for _, bsLine := range order.Lines {
		lines = append(lines, ToOrderLineModel(bsLine))
	}
	return models.OrderModel{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		CreatedAt:  order.CreatedAt,
		Status:     string(order.Status),
		Lines:      lines,
	}
}

func ToOrder(model models.OrderModel) domain.Order {
	var lines []domain.OrderLine
	for _, modelLine := range model.Lines {
		lines = append(lines, ToOrderLine(modelLine))
	}
	return domain.Order{
		ID:         model.ID,
		CustomerID: model.CustomerID,
		CreatedAt:  model.CreatedAt,
		Status:     domain.OrderStatus(model.Status),
		Lines:      lines,
	}
}

// orderline mappers
func ToOrderLineModel(orderL domain.OrderLine) models.OrderLineModel {
	return models.OrderLineModel{
		ID:         orderL.ID,
		OrderID:    orderL.OrderID,
		ProductID:  orderL.ProductID,
		LocationID: orderL.LocationID,
		Quantity:   orderL.Quantity,
	}
}

func ToOrderLine(model models.OrderLineModel) domain.OrderLine {
	return domain.OrderLine{
		ID:         model.ID,
		OrderID:    model.OrderID,
		ProductID:  model.ProductID,
		LocationID: model.LocationID,
		Quantity:   model.Quantity,
	}
}
