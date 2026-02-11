package usecase

import (
	"context"

	"github.com/trng-tr/order-microservice/internal/application/out"
	"github.com/trng-tr/order-microservice/internal/domain"
)

type RemoteCustomerServiceImpl struct {
	outSvc out.RemoteCustomerService
}

func NewRemoteCustomerServiceImpl(outS out.RemoteCustomerService) *RemoteCustomerServiceImpl {
	return &RemoteCustomerServiceImpl{outSvc: outS}
}

func (o *RemoteCustomerServiceImpl) GetRemoteCustomerByID(ctx context.Context, id int64) (domain.Customer, error) {
	/*if err := checkId(id); err != nil {
		return domain.Customer{}, err
	}*/

	bsCustomer, err := o.outSvc.GetRemoteCustomerByID(ctx, id)
	if err != nil {
		return domain.Customer{}, err
	}

	return bsCustomer, nil

}
