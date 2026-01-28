package usecase

import (
	"context"
	"fmt"

	"github.com/trng-tr/customer-microservice/internal/application/ports/out"
	"github.com/trng-tr/customer-microservice/internal/domain"
)

/*
InCustomerServiceImpl implement port d'entrée exposé à l'extreieur
il utilise pour cela le port de sortie OutCustomerService et le port
de sortie OutAddressService pour récuprer l'adresse
*/
type InCustomerServiceImpl struct {
	outCustomerSvc out.OutCustomerService
	outAddressSvc  out.OutAddressService
}

// NewInCustomerServiceImpl DI by constructor
func NewInCustomerServiceImpl(outCustomerSvc out.OutCustomerService, outAddressSvc out.OutAddressService) *InCustomerServiceImpl {
	return &InCustomerServiceImpl{outCustomerSvc: outCustomerSvc, outAddressSvc: outAddressSvc}
}

// CreateCustomer implement interface InCustomerService
func (icsi *InCustomerServiceImpl) CreateCustomer(ctx context.Context, bsCustomer domain.Customer) (domain.Customer, error) {
	var inputFields map[string]string = map[string]string{
		"firstname": bsCustomer.Firstname,
		"lastname":  bsCustomer.Lastname,
		"email":     bsCustomer.Email,
		"phone":     bsCustomer.PhoneNumber,
	}
	if err := checkInputFields(inputFields); err != nil {
		return domain.Customer{}, err
	}
	if err := checkInputGenda(bsCustomer.Genda); err != nil {
		return domain.Customer{}, err
	}
	if ok := checkEmailValid(bsCustomer.Email); !ok {
		return domain.Customer{}, fmt.Errorf("error: invalid input email %s", bsCustomer.Email)
	}
	if err := checkPhoneValid(bsCustomer.PhoneNumber); err != nil {
		return domain.Customer{}, err
	}

	bsCustomer.CreatedAt = generateDate()
	bsCustomer.Status = domain.Active
	//check address
	if _, err := icsi.outAddressSvc.GetAddressByID(ctx, bsCustomer.AddressID); err != nil {
		return domain.Customer{}, err
	}
	// send business object to outside using output port
	sentBsCustomer, err := icsi.outCustomerSvc.CreateCustomer(ctx, bsCustomer)
	if err != nil {
		return domain.Customer{}, fmt.Errorf("save customer failed %w", err)
	}

	return sentBsCustomer, nil
}

// GetCustomerByID implement interface InCustomerService
func (icsi *InCustomerServiceImpl) GetCustomerByID(ctx context.Context, id int64) (domain.Customer, error) {
	if err := checkInputId(id); err != nil {
		return domain.Customer{}, err
	}

	bsCustomer, err := icsi.outCustomerSvc.GetCustomerByID(ctx, id)
	if err != nil {
		return domain.Customer{}, err
	}

	return bsCustomer, nil

}

// GetAllCustomers implement interface InCustomerService
func (icsi *InCustomerServiceImpl) GetAllCustomers(ctx context.Context) ([]domain.Customer, error) {
	// call output port to retrieve all customers
	return icsi.outCustomerSvc.GetAllCustomers(ctx)
}

// UpdateCustomer implement interface InCustomerService
func (icsi *InCustomerServiceImpl) PatchCustomer(ctx context.Context, id int64, patchCustomer domain.PatchBusinessCustomer) (domain.Customer, error) {
	if err := checkInputId(id); err != nil {
		return domain.Customer{}, err
	}
	businessCustomer, err := icsi.outCustomerSvc.GetCustomerByID(ctx, id)
	if err != nil {
		return domain.Customer{}, err
	}
	businessCustomer.ApplyPatchCustomer(patchCustomer) //mapper for patch request

	// call outputservice to save changes
	sentUpdatedCustomer, err := icsi.outCustomerSvc.PatchCustomer(ctx, id, businessCustomer)
	if err != nil {
		return domain.Customer{}, err
	}

	return sentUpdatedCustomer, nil
}

// DeleteCustomer implement interface InCustomerService
func (icsi *InCustomerServiceImpl) DeleteCustomer(ctx context.Context, id int64) error {
	if err := checkInputId(id); err != nil {
		return err
	}

	if err := icsi.outCustomerSvc.DeleteCustomer(ctx, id); err != nil {
		return err
	}

	return nil
}
