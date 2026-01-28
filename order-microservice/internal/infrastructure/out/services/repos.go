package services

import (
	"context"

	"github.com/trng-tr/order-microservice/internal/infrastructure/out/models"
)

// OrderRepo extends Repository
type OrderRepo interface {
	Save(ctx context.Context, orderModel models.OrderModel) (models.OrderModel, error)
	Repository[models.OrderModel, int64]
}

// OrderLineRepo extends Repository
type OrderLineRepo interface {
	Repository[models.OrderLineModel, int64]
	Update(ctx context.Context, id int64, quantity int64) (models.OrderLineModel, error)
	FindAllByOrderID(ctx context.Context, orderID int64) ([]models.OrderLineModel, error)
}
