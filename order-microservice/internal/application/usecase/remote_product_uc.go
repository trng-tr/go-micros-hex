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
