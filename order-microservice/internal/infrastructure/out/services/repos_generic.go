package services

import "context"

type Repository[O any, ID comparable] interface {
	FindByID(ctx context.Context, id ID) (O, error)
	FindAll(ctx context.Context) ([]O, error)
	Delete(ctx context.Context, id ID) error
}
