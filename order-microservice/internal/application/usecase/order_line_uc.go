package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/trng-tr/order-microservice/internal/application/out"
	"github.com/trng-tr/order-microservice/internal/domain"
)

// OrderLineUseCase implement OrderLineService
type OrderLineUseCase struct {
	outOrderLineSvc out.OutOrderLineService  //DI
	outOrderSvc     out.OutOrderService      //DI
	remoteProduct   out.RemoteProductService //DI
	remoteStock     out.RemoteStockService   //DI
}

// NewOrderLineServiceImpl DI par constructeur
func NewOrderLineUseCase(outOrderLineSvc out.OutOrderLineService, outOrderSvc out.OutOrderService,
	remoteProduct out.RemoteProductService, remoteStock out.RemoteStockService) *OrderLineUseCase {
	return &OrderLineUseCase{
		outOrderLineSvc: outOrderLineSvc,
		outOrderSvc:     outOrderSvc,
		remoteProduct:   remoteProduct,
		remoteStock:     remoteStock,
	}
}

// GetOrderByID implement OrderLineService interface
func (o *OrderLineUseCase) GetOrderLineByID(ctx context.Context, id int64) (domain.OrderLine, error) {
	savedOrder, err := o.outOrderLineSvc.GetOrderLineByID(ctx, id)
	if err != nil {
		return domain.OrderLine{}, fmt.Errorf("%w:%v", errOccurred, err)
	}
	return savedOrder, nil
}

// GetAllOrder implement OrderLineService interface
func (o *OrderLineUseCase) GetAllOrderLines(ctx context.Context) ([]domain.OrderLine, error) {
	orderlines, err := o.outOrderLineSvc.GetAllOrderLines(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w:%v", errOccurred, err)
	}

	return orderlines, nil
}

// SetOrderLineQuantity implement OrderLineService interface
func (o *OrderLineUseCase) SetOrderLineQuantity(ctx context.Context, id int64, quantity int64) (domain.OrderLine, error) {
	values := map[string]int64{
		"id":       id,
		"quantity": quantity,
	}
	if err := checkValue(values); err != nil {
		return domain.OrderLine{}, err
	}
	savedOrderLine, err := o.GetOrderLineByID(ctx, id)
	if err != nil {
		return domain.OrderLine{}, fmt.Errorf("%w:%v", errOccurred, err)
	}
	savedOrderLine.Quantity = quantity

	UpdateOrderLine, err := o.outOrderLineSvc.UpdateOrderLine(ctx, savedOrderLine)
	if err != nil {
		return domain.OrderLine{}, fmt.Errorf("%w:%v", errOccurred, err)
	}

	return UpdateOrderLine, nil
}

// IncreaseOrderLineQuantity implement OrderLineService interface
func (o *OrderLineUseCase) IncreaseOrderLineQuantity(ctx context.Context, id int64, quantity int64) (domain.OrderLine, error) {
	values := map[string]int64{
		"id":       id,
		"quantity": quantity,
	}
	if err := checkValue(values); err != nil {
		return domain.OrderLine{}, err
	}
	savedLine, err := o.GetOrderLineByID(ctx, id)
	if err != nil {
		return domain.OrderLine{}, fmt.Errorf("%w:%v", errOccurred, err)
	}
	//check remote product exist again and is active
	product, err := o.remoteProduct.GetRemoteProductByID(ctx, savedLine.ProductID)
	if err != nil {
		return domain.OrderLine{}, err
	}
	if ok := product.IsActive; !ok {
		return domain.OrderLine{}, errors.New("error: remote product status not allowed")
	}
	//check if there is enough quantity in stock
	stock, err := o.remoteStock.GetRemoteStockByLocationIDAndProductID(ctx, savedLine.LocationID, savedLine.ProductID)
	if err != nil {
		return domain.OrderLine{}, err
	}
	if (stock.Quantity - quantity) < 0 {
		return domain.OrderLine{}, fmt.Errorf("%w for product %d", errNotEnough, stock.ProductID)
	}
	savedLine.Quantity += quantity
	UpdateOrderLine, err := o.outOrderLineSvc.UpdateOrderLine(ctx, savedLine)
	if err != nil {
		return domain.OrderLine{}, fmt.Errorf("%w:%v", errOccurred, err)
	}
	//prendre du stock la quantité augmentée à la ligne de commande
	stock.Quantity -= quantity
	// call remote to update stock
	if err := o.remoteStock.SetRemoteStockQuantity(ctx, stock.LocationID, stock.ProductID, stock); err != nil {
		return domain.OrderLine{}, err
	}
	return UpdateOrderLine, nil
}

// DecreaseOrderLineQuantity implement OrderLineService interface
func (o *OrderLineUseCase) DecreaseOrderLineQuantity(ctx context.Context, id int64, quantity int64) (domain.OrderLine, error) {
	values := map[string]int64{
		"id":       id,
		"quantity": quantity,
	}
	if err := checkValue(values); err != nil {
		return domain.OrderLine{}, err
	}
	savedLine, err := o.GetOrderLineByID(ctx, id)
	if err != nil {
		return domain.OrderLine{}, fmt.Errorf("%w:%v", errOccurred, err)
	}

	//check remote product exist again
	// for decrease quantity, no need to check if prodct is active or not
	_, err = o.remoteProduct.GetRemoteProductByID(ctx, savedLine.ProductID)
	if err != nil {
		return domain.OrderLine{}, err
	}
	//check stock is recheable
	stock, err := o.remoteStock.GetRemoteStockByLocationIDAndProductID(ctx, savedLine.LocationID, savedLine.ProductID)
	if err != nil {
		return domain.OrderLine{}, err
	}
	if savedLine.Quantity <= quantity {
		return domain.OrderLine{}, errors.New("error: quantity to decrease exceeds current order line quantity")
	}
	savedLine.Quantity -= quantity
	UpdateOrderLine, err := o.outOrderLineSvc.UpdateOrderLine(ctx, savedLine)
	if err != nil {
		return domain.OrderLine{}, fmt.Errorf("%w:%v", errOccurred, err)
	}
	//remettre en stock la quantité diminiuée de la ligne de commande
	stock.Quantity += quantity
	if err := o.remoteStock.SetRemoteStockQuantity(ctx, stock.LocationID, stock.ProductID, stock); err != nil {
		return domain.OrderLine{}, err
	}

	return UpdateOrderLine, nil
}

// DeleteOrderLine implement OrderLineService interface
func (o *OrderLineUseCase) DeleteOrderLine(ctx context.Context, id int64) error {
	_, err := o.GetOrderLineByID(ctx, id)
	if err != nil {
		return fmt.Errorf("%w:%v", errOccurred, err)
	}

	if err := o.outOrderLineSvc.DeleteOrderLine(ctx, id); err != nil {
		return fmt.Errorf("%w:%v", errOccurred, err)
	}

	return nil
}

// GetOrderLinesByOrderID implement OrderLineService interface
func (o *OrderLineUseCase) GetOrderLinesByOrderID(ctx context.Context, orderID int64) ([]domain.OrderLine, error) {
	orderLines, err := o.outOrderLineSvc.GetOrderLinesByOrderID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("%w:%v", errOccurred, err)
	}

	return orderLines, nil
}
