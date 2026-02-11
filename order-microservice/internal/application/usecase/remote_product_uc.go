package usecase

import (
	"context"
	"fmt"

	"github.com/trng-tr/order-microservice/internal/application/out"
	"github.com/trng-tr/order-microservice/internal/domain"
)

// RemoteProductServiceImpl implement interface
type RemoteProductServiceImpl struct {
	outSvc out.RemoteProductService
}

// NewRemoteProductServiceImpl DI par constructeur
func NewRemoteProductServiceImpl(outS out.RemoteProductService) *RemoteProductServiceImpl {
	return &RemoteProductServiceImpl{outSvc: outS}
}

// GetRemoteProductByID immplement interface
func (o *RemoteProductServiceImpl) GetRemoteProductByID(ctx context.Context, productID int64) (domain.Product, error) {
	inputs := map[string]int64{
		"product_id": productID,
	}
	if err := checkValue(inputs); err != nil {
		return domain.Product{}, err
	}

	bsProduct, err := o.outSvc.GetRemoteProductByID(ctx, productID)
	if err != nil {
		return domain.Product{}, fmt.Errorf("%w:%v", errOccurred, err)
	}

	return bsProduct, nil
}

// GetRemoteStockByProductID immplement interface
func (o *RemoteProductServiceImpl) GetRemoteStockByLocationIDAndProductID(ctx context.Context, locationID, prodID int64) (domain.Stock, error) {
	values := map[string]int64{
		"location_id": locationID,
		"product_id":  prodID,
	}
	if err := checkValue(values); err != nil {
		return domain.Stock{}, err
	}
	stock, err := o.outSvc.GetRemoteStockByLocationIDAndProductID(ctx, locationID, prodID)
	if err != nil {
		return domain.Stock{}, fmt.Errorf("%w:%v", errOccurred, err)
	}

	return stock, nil
}

// SetRemoteStockQuantity immplement interface
func (o *RemoteProductServiceImpl) SetRemoteStockQuantity(ctx context.Context, productID, locationID, newQuantity int64) error {
	values := map[string]int64{
		"product_id":  productID,
		"location_id": locationID,
		"quantity":    newQuantity,
	}
	if err := checkValue(values); err != nil {
		return err
	}

	stock, err := o.outSvc.GetRemoteStockByLocationIDAndProductID(ctx, productID, locationID)
	if err != nil {
		return fmt.Errorf("%w:%v", errOccurred, err)
	}
	stock.Quantity -= newQuantity
	// call remote service to send for update remote stock
	if err := o.outSvc.SetRemoteStockQuantity(ctx, productID, locationID, stock); err != nil {
		return fmt.Errorf("%w:%v", errOccurred, err)
	}

	return nil
}
