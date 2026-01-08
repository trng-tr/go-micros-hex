package out

import (
	"context"

	"github.com/trng-tr/customer-microservice/internal/domain"
)

// OutAddressService port de sortie utilisé par l'app pour envoyer à l'exterieur
type OutAddressService interface {
	CreateAddress(ctx context.Context, address domain.BusinessAddress) (domain.BusinessAddress, error)
	GetAddressByID(ctx context.Context, id int64) (domain.BusinessAddress, error)
}
