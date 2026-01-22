package services

import (
	"context"

	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/out/mappers"
)

// OutOrderServiceImpl implement interface OutOrderService
type OutOrderServiceImpl struct {
	repo OrderRepo
}

// NewOutOrderServiceImpl DI by constructor
func NewOutOrderServiceImpl(repo OrderRepo) *OutOrderServiceImpl {
	return &OutOrderServiceImpl{repo: repo}
}

// CreateOrder implement interface OutOrderService
func (o *OutOrderServiceImpl) CreateOrderWithOrderLines(ctx context.Context, order domain.Order) (domain.Order, error) {
	model := mappers.ToOrderModel(order)
	saved, err := o.repo.Save(ctx, model)
	if err != nil {
		return domain.Order{}, err
	}

	return mappers.ToOrder(saved), nil
}

// GetOrderByID implement interface OutOrderService
func (o *OutOrderServiceImpl) GetOrderByID(ctx context.Context, id int64) (domain.Order, error) {
	model, err := o.repo.FindByID(ctx, id)
	if err != nil {
		return domain.Order{}, err
	}

	return mappers.ToOrder(model), nil
}

// GetAllOrder implement interface OutOrderService
func (o *OutOrderServiceImpl) GetAllOrder(ctx context.Context) ([]domain.Order, error) {
	models, err := o.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var orders = make([]domain.Order, 0, len(models))
	for _, model := range models {
		orders = append(orders, mappers.ToOrder(model))
	}

	return orders, nil
}

// DeleteOrder implement interface OutOrderService
func (o *OutOrderServiceImpl) DeleteOrder(ctx context.Context, id int64) error {
	if err := o.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
