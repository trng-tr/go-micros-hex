package out

import (
	"context"

	"github.com/trng-tr/order-microservice/internal/domain"
)

type OutOrderService interface {
	CreateOrderWithOrderLines(ctx context.Context, order domain.Order) (domain.Order, error)
	GetOrderByID(ctx context.Context, id int64) (domain.Order, error)
	GetAllOrders(ctx context.Context) ([]domain.Order, error)
	DeleteOrder(ctx context.Context, id int64) error
}

type OutOrderLineService interface {
	GetOrderLineByID(ctx context.Context, id int64) (domain.OrderLine, error)
	GetAllOrderLines(ctx context.Context) ([]domain.OrderLine, error)
	UpdateOrderLine(ctx context.Context, orderLine domain.OrderLine) (domain.OrderLine, error)
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
	SetRemoteStockQuantity(ctx context.Context, productID, locationID int64, stock domain.Stock) error
}

type RemoteStockService interface {
	GetRemoteStockByID(ctx context.Context, stockID int64) (domain.Location, error)
}

type RemoteLocationService interface {
	GetRemoteLocationByID(ctx context.Context, locationID int64) (domain.Location, error)
}
