package in

import (
	"context"

	"github.com/trng-tr/order-microservice/internal/domain"
)

type InOrderService interface {
	CreateOrderWithOrderLines(ctx context.Context, customerID int64, lines []domain.OrderLine) (domain.Order, error)
	GetOrderByID(ctx context.Context, id int64) (domain.Order, error)
	GetAllOrders(ctx context.Context) ([]domain.Order, error)
	DeleteOrder(ctx context.Context, id int64) error
}
type InOrderLineService interface {
	GetOrderLineByID(ctx context.Context, id int64) (domain.OrderLine, error)
	GetAllOrderLines(ctx context.Context) ([]domain.OrderLine, error)
	SetOrderLineQuantity(ctx context.Context, id int64, quantity int64) (domain.OrderLine, error)
	IncreaseOrderLineQuantity(ctx context.Context, id int64, quantity int64) (domain.OrderLine, error)
	DecreaseOrderLineQuantity(ctx context.Context, id int64, quantity int64) (domain.OrderLine, error)
	DeleteOrderLine(ctx context.Context, id int64) error
	GetOrderLinesByOrderID(ctx context.Context, orderID int64) ([]domain.OrderLine, error)
}

// RemoteCustomerService to get remote customer
type RemoteCustomerService interface {
	GetRemoteCustomerByID(ctx context.Context, id int64) (domain.Customer, error)
}

// RemoteProductService to get remote products
type RemoteProductService interface {
	GetRemoteProductByID(ctx context.Context, productID int64) (domain.Product, error)
	GetRemoteStockByLocationIDAndProductID(ctx context.Context, locationID, prodID int64) (domain.Stock, error)
	SetRemoteStockQuantity(ctx context.Context, productID, locationID, newQuantity int64) error
}

type RemoteStockService interface {
	GetRemoteStockByID(ctx context.Context, stockID int64) (domain.Location, error)
}

type RemoteLocationService interface {
	GetRemoteLocationByID(ctx context.Context, locationID int64) (domain.Location, error)
}
