package out

import (
	"context"

	"github.com/trng-tr/customer-microservice/internal/domain"
)

// OutCustomerService port de sortie utilisé par l'app pour envoyer à l'exterieur
type OutCustomerService interface {
	CreateCustomer(ctx context.Context, customer domain.Customer) (domain.Customer, error)
	GetCustomerByID(ctx context.Context, id int64) (domain.Customer, error)
	GetAllCustomers(ctx context.Context) ([]domain.Customer, error)
	PatchCustomer(ctx context.Context, id int64, businessCustomer domain.Customer) (domain.Customer, error)
	DeleteCustomer(ctx context.Context, id int64) error
}
