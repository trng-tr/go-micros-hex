package services

import (
	"context"
	"fmt"

	"github.com/trng-tr/customer-microservice/internal/domain"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/mappers"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/models"
)

/*
OutCustomerServiceImpl implementation of output ports in infra
tp persit in db, we inject customer repository
*/
type OutCustomerServiceImpl struct {
	repo CustomerRepository
}

// NewOutCustomerServiceImpl DI by constructor
func NewOutCustomerServiceImpl(repo CustomerRepository) *OutCustomerServiceImpl {
	return &OutCustomerServiceImpl{repo: repo}
}

// CreateCustomer implement OutCustomerService interface
func (ocsi *OutCustomerServiceImpl) CreateCustomer(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	var customerModel models.CustomerModel = mappers.ToCustomerModel(customer)
	savedCustomer, err := ocsi.repo.SaveO(ctx, customerModel)
	if err != nil {
		return domain.Customer{}, err
	}

	return mappers.ToBusinessCustomer(savedCustomer), nil
}

// GetCustomerByID implement OutCustomerService interface
func (ocsi *OutCustomerServiceImpl) GetCustomerByID(ctx context.Context, id int64) (domain.Customer, error) {
	savedCustomer, err := ocsi.repo.FindOByID(ctx, id)
	if err != nil {
		return domain.Customer{}, fmt.Errorf("find customer by id failed %w", err) //%w pour wrapper err
	}
	return mappers.ToBusinessCustomer(savedCustomer), nil
}

// GetAllCustomers implement OutCustomerService interface
func (ocsi *OutCustomerServiceImpl) GetAllCustomers(ctx context.Context) ([]domain.Customer, error) {
	modelsCustomers, err := ocsi.repo.FindAllO(ctx)
	if err != nil {
		return nil, fmt.Errorf("find all customer failed %w", err) //%w pour wrapper err
	}

	var bsCustomers []domain.Customer = make([]domain.Customer, 0, len(modelsCustomers))
	for _, model := range modelsCustomers {
		bsCustomers = append(bsCustomers, mappers.ToBusinessCustomer(model))
	}

	return bsCustomers, nil

}

// PatchCustomer implement OutCustomerService interface
func (ocsi *OutCustomerServiceImpl) PatchCustomer(ctx context.Context, id int64, businessCustomer domain.Customer) (domain.Customer, error) {
	var model models.CustomerModel = mappers.ToCustomerModel(businessCustomer)
	patchedCustomerModel, err := ocsi.repo.UpdateO(ctx, id, model)
	if err != nil {
		return domain.Customer{}, err
	}
	return mappers.ToBusinessCustomer(patchedCustomerModel), nil
}

func (ocsi *OutCustomerServiceImpl) DeleteCustomer(ctx context.Context, id int64) error {
	if err := ocsi.repo.DeleteO(ctx, id); err != nil {
		return err
	}
	return nil
}
