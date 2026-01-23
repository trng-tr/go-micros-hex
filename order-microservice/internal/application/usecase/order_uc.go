package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/trng-tr/order-microservice/internal/application/out"
	"github.com/trng-tr/order-microservice/internal/domain"
)

// OrderUseCase implement OrderService
type OrderUseCase struct {
	outOrderSvc       out.OutOrderService
	remoteCustomerSvc out.RemoteCustomerService
	remoteProductSvc  out.RemoteProductService
}

// NewOrderNewOrderUseCaseServiceImpl DI by contructor
func NewOrderUseCase(outOrderSvc out.OutOrderService, remote1 out.RemoteCustomerService,
	remote2 out.RemoteProductService) *OrderUseCase {
	return &OrderUseCase{outOrderSvc: outOrderSvc, remoteCustomerSvc: remote1, remoteProductSvc: remote2}
}

// CreateOrder implement OrderService
func (o *OrderUseCase) CreateOrderWithOrderLines(ctx context.Context, customerID int64, lines []domain.OrderLine) (domain.Order, error) {
	if err := checkId(customerID); err != nil {
		return domain.Order{}, err
	}
	if len(lines) == 0 {
		return domain.Order{}, errors.New("error: order must contain lines")
	}
	for i := range lines {
		values := map[string]int64{
			"product_id": lines[i].ProductID,
			"quantity":   lines[i].Quantity,
		}
		if err := checkValue(values); err != nil {
			return domain.Order{}, fmt.Errorf("%w:%v", errOccurred, err)
		}
	}
	// call remote service to check remote customer 👇
	customer, err := o.remoteCustomerSvc.GetRemoteCustomerByID(ctx, customerID)
	if err != nil {
		return domain.Order{}, fmt.Errorf("%w:%v", errOccurred, err)
	}
	if customer.Status != domain.Active {
		return domain.Order{}, errors.New("error: remote customer status not allowed")
	}

	// 3) remote product + stock checks (best-effort)
	// ⚠️ Ici on vérifie juste avant commit.
	// Pour du 100% robuste, il faut réserver le stock (saga).👇
	var stocksToUpdate []domain.Stock = make([]domain.Stock, 0, len(lines))
	for _, line := range lines {
		remoteProduct, err := o.remoteProductSvc.GetRemoteProductByID(ctx, line.ProductID)
		if err != nil {
			return domain.Order{}, fmt.Errorf("%w:%v", errOccurred, err)
		}
		if ok := remoteProduct.IsActive; !ok {
			return domain.Order{}, errors.New("error: remote product status not allowed")
		}
		// get stock for the product to check quantity is enough 👇
		stock, err := o.remoteProductSvc.GetRemoteStockByProductID(ctx, line.ProductID)
		if err != nil {
			return domain.Order{}, fmt.Errorf("%w:%v", errOccurred, err)
		}
		if (stock.Quantity - line.Quantity) < 0 {
			return domain.Order{}, fmt.Errorf("%w for product %d", errNotEnough, stock.ProductID)
		}
		stock.Quantity -= line.Quantity
		stocksToUpdate = append(stocksToUpdate, stock)
	}

	//4 build order object to send to output service
	var order = domain.Order{
		CustomerID: customerID,
		CreatedAt:  time.Now(),
		Status:     domain.Created,
		Lines:      lines,
	}

	// 5) atomic DB: create order + all lines (transaction inside outOrderSvc)
	savedOrder, err := o.outOrderSvc.CreateOrderWithOrderLines(ctx, order)
	if err != nil {
		return domain.Order{}, fmt.Errorf("%w:%v", errOccurred, err)
	}

	// 6) update remote stock AFTER DB commit (best-effort)
	// ⚠️ Si ça échoue, tu dois compenser (annuler commande) ou marquer FAILED.👇
	for _, stock := range stocksToUpdate {
		if err := o.remoteProductSvc.SetRemoteStockQuantity(ctx, stock.ProductID, stock); err != nil {
			return domain.Order{}, fmt.Errorf("%w:%v", errOccurred, err)
		}
	}

	return savedOrder, nil

}

// GetOrderByID implement OrderService
func (o *OrderUseCase) GetOrderByID(ctx context.Context, id int64) (domain.Order, error) {
	if err := checkId(id); err != nil {
		return domain.Order{}, err
	}

	savedOrder, err := o.outOrderSvc.GetOrderByID(ctx, id)
	if err != nil {
		return domain.Order{}, fmt.Errorf("%w:%v", errOccurred, err)
	}

	return savedOrder, nil
}

// GetAllOrder implement OrderService
func (o *OrderUseCase) GetAllOrder(ctx context.Context) ([]domain.Order, error) {
	orders, err := o.outOrderSvc.GetAllOrder(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w:%v", errOccurred, err)
	}
	if len(orders) == 0 {
		return nil, fmt.Errorf("%w", errNoRows)
	}

	return orders, nil
}

// DeleteOrder implement OrderService
func (o *OrderUseCase) DeleteOrder(ctx context.Context, id int64) error {
	if err := checkId(id); err != nil {
		return err
	}
	order, err := o.GetOrderByID(ctx, id)
	if err != nil {
		return fmt.Errorf("%w:%v", errOccurred, err)
	}
	if order.Status != domain.Created {
		return errors.New("order can no longer be deleted")
	}
	if err := o.outOrderSvc.DeleteOrder(ctx, id); err != nil {
		return fmt.Errorf("%w:%v", errOccurred, err)
	}

	return nil
}
