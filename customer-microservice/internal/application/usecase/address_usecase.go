package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/trng-tr/customer-microservice/internal/application/ports/out"
	"github.com/trng-tr/customer-microservice/internal/domain"
)

/*
AddressServiceImpl implement port d'entrée exposé à l'extreieur
il utilise pour cela le port de sortie OutAddressService
*/
type InAddressServiceImpl struct {
	outService out.OutAddressService
}

// NewAddressServiceImpl DI by constructor
func NewInAddressServiceImpl(outService out.OutAddressService) *InAddressServiceImpl {
	return &InAddressServiceImpl{outService: outService}
}

// CreateAddress implement InAddressService interface
func (iasi *InAddressServiceImpl) CreateAddress(ctx context.Context, bsAddress domain.Address) (domain.Address, error) {
	var inputFields map[string]string = map[string]string{
		"street_name": bsAddress.StreetName,
		"zip_code":    bsAddress.ZipCode,
		"city":        bsAddress.City,
		"region":      bsAddress.Region,
		"country":     bsAddress.Country,
	}
	var err error = checkInputFields(inputFields)
	if err != nil {
		return domain.Address{}, err
	}
	if err := checkZipCode(bsAddress.ZipCode); err != nil {
		return domain.Address{}, err
	}
	// send business object to outside using output port
	sendBsAddress, err := iasi.outService.CreateAddress(ctx, bsAddress)
	if err != nil {
		return domain.Address{}, fmt.Errorf("%w:%v", errSendObject, err)
	}

	return sendBsAddress, nil
}

// GetAddressByID implement InAddressService interface
func (iasi *InAddressServiceImpl) GetAddressByID(ctx context.Context, id int64) (domain.Address, error) {
	if err := checkInputId(id); err != nil {
		return domain.Address{}, err
	}
	//call output port to retried business address object
	bsAddress, err := iasi.outService.GetAddressByID(ctx, id)
	if err != nil {
		return domain.Address{}, fmt.Errorf("%w:%v", errRetrieveObject, err)
	}

	return bsAddress, nil
}

// GetAllAddresses implement InAddressService interface
func (iasi *InAddressServiceImpl) GetAllAddresses(ctx context.Context) ([]domain.Address, error) {
	addresses, err := iasi.outService.GetAllAddresses(ctx)
	if err != nil {
		return nil, fmt.Errorf("an arror has occurred %w", err)
	}
	if len(addresses) == 0 {
		return nil, errors.New("no address is registered")
	}

	return addresses, nil
}

// DeleteAddress implement InAddressService interface
func (iasi *InAddressServiceImpl) DeleteAddress(ctx context.Context, id int64) error {
	if err := checkInputId(id); err != nil {
		return fmt.Errorf("error: delete address failed %w", err)
	}
	if err := iasi.outService.DeleteAddress(ctx, id); err != nil {
		return fmt.Errorf("error: delete address failed %w", err)
	}

	return nil
}
