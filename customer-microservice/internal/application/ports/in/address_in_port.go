package in

import (
	"context"

	"github.com/trng-tr/customer-microservice/internal/domain"
)

// AddressService port d'entrée exposé par l'app à l'exterieur
type InAddressService interface {
	CreateAddress(ctx context.Context, bsAddress domain.Address) (domain.Address, error)
	GetAddressByID(ctx context.Context, id int64) (domain.Address, error)
	GetAllAddresses(ctx context.Context) ([]domain.Address, error)
	DeleteAddress(ctx context.Context, id int64) error
}
