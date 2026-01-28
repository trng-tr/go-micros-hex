package services

import (
	"context"

	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/out/mappers"
	"github.com/trng-tr/product-microservice/internal/infrastructure/out/models"
)

// OutLocationServiceImpl implements interface OutLocationService
type OutLocationServiceImpl struct {
	repo Repository[models.Location, int64] //DI
}

// NewOutLocationServiceImpl DI by constructeur
func NewOutLocationServiceImpl(repo Repository[models.Location, int64]) *OutLocationServiceImpl {
	return &OutLocationServiceImpl{repo: repo}
}

// CreateLocation implements interface OutLocationService
func (o *OutLocationServiceImpl) CreateLocation(ctx context.Context, localtion domain.Location) (domain.Location, error) {
	model, err := o.repo.Save(ctx, mappers.ToLocationModel(localtion))
	if err != nil {
		return domain.Location{}, err
	}

	return mappers.ToBusinessLocation(model), nil
}

// GetLocationByID implements interface OutLocationService
func (o *OutLocationServiceImpl) GetLocationByID(ctx context.Context, id int64) (domain.Location, error) {
	model, err := o.repo.FindByID(ctx, id)
	if err != nil {
		return domain.Location{}, err
	}

	return mappers.ToBusinessLocation(model), nil
}

func (o *OutLocationServiceImpl) GetAllLocation(ctx context.Context) ([]domain.Location, error) {
	models, err := o.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var bsLocations []domain.Location = make([]domain.Location, len(models))
	for i, m := range models {
		bsLocations[i] = mappers.ToBusinessLocation(m)
	}

	return bsLocations, nil
}
