package out

import (
	"context"

	"github.com/trng-tr/customer-microservice/internal/domain"
)

// OutCustomerService port de sortie utilisé par l'app pour envoyer à l'exterieur
type OutCustomerService interface {
	CreateCustomer(ctx context.Context, customer domain.BusinessCustomer) (domain.BusinessCustomer, error)
	GetCustomerByID(ctx context.Context, id int64) (domain.BusinessCustomer, error)
	GetAllCustomers(ctx context.Context) ([]domain.BusinessCustomer, error)
	UpdateCustomer(ctx context.Context, id int64, customer domain.BusinessCustomer) (domain.BusinessCustomer, error)
	DeleteCustomer(ctx context.Context, id int64) error
}
