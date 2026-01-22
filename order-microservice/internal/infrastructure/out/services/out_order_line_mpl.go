package services

import (
	"context"

	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/out/mappers"
)

// OutOrderLineServiceImpl implement interface OutOrderLineService
type OutOrderLineServiceImpl struct {
	repo OrderLineRepo
}

// NewOutOrderLineServiceImpl DI by constrcutor
func NewOutOrderLineServiceImpl(repo OrderLineRepo) *OutOrderLineServiceImpl {
	return &OutOrderLineServiceImpl{repo: repo}
}

// GetOrderLineByID implements OutOrderLineService
func (o *OutOrderLineServiceImpl) GetOrderLineByID(ctx context.Context, id int64) (domain.OrderLine, error) {
	model, err := o.repo.FindByID(ctx, id)
	if err != nil {
		return domain.OrderLine{}, err
	}
	return mappers.ToOrderLine(model), nil
}

// GetAllOrderLines implements OutOrderLineService
func (o *OutOrderLineServiceImpl) GetAllOrderLines(ctx context.Context) ([]domain.OrderLine, error) {
	models, err := o.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	orderLines := make([]domain.OrderLine, 0, len(models))
	for _, model := range models {
		orderLines = append(orderLines, mappers.ToOrderLine(model))
	}

	return orderLines, nil
}

// UpdateOrderLine implements OutOrderLineService
func (o *OutOrderLineServiceImpl) UpdateOrderLine(ctx context.Context, orderLine domain.OrderLine) (domain.OrderLine, error) {
	model := mappers.ToOrderLineModel(orderLine)
	updatedOrderLine, err := o.repo.Update(ctx, orderLine.ID, model.Quantity)
	if err != nil {
		return domain.OrderLine{}, err
	}

	return mappers.ToOrderLine(updatedOrderLine), nil
}

// DeleteOrderLine implements OutOrderLineService
func (o *OutOrderLineServiceImpl) DeleteOrderLine(ctx context.Context, id int64) error {
	if err := o.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

// GetOrderLinesByOrderID implement OrderLineService interface
func (o *OutOrderLineServiceImpl) GetOrderLinesByOrderID(ctx context.Context, orderID int64) ([]domain.OrderLine, error) {
	models, err := o.repo.FindAllByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	var orderLines []domain.OrderLine
	for _, model := range models {
		orderLines = append(orderLines, mappers.ToOrderLine(model))
	}

	return orderLines, nil
}
