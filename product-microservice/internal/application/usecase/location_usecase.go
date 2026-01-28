package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/trng-tr/product-microservice/internal/application/ports/out"
	"github.com/trng-tr/product-microservice/internal/domain"
)

// InLocationServiceImpl implements interface InLocationService
type InLocationServiceImpl struct {
	output out.OutLocationService
}

// NewInLocationServiceImpl DI par constructor
func NewInLocationServiceImpl(output out.OutLocationService) *InLocationServiceImpl {
	return &InLocationServiceImpl{output: output}
}

// CreateLocation implement InLocationService interface
func (i *InLocationServiceImpl) CreateLocation(ctx context.Context, location domain.Location) (domain.Location, error) {
	inputs := map[string]string{
		"city": location.Ville,
	}
	if err := checkInputs(inputs); err != nil {
		return domain.Location{}, err
	}

	location.CreatedAt = time.Now()

	savedLocation, err := i.output.CreateLocation(ctx, location)
	if err != nil {
		return domain.Location{}, fmt.Errorf("%w", err)
	}

	return savedLocation, nil
}

// GetLocationByID implement InLocationService interface
func (i *InLocationServiceImpl) GetLocationByID(ctx context.Context, id int64) (domain.Location, error) {
	if err := checkInputId(id); err != nil {
		return domain.Location{}, err
	}

	savedLocation, err := i.output.GetLocationByID(ctx, id)
	if err != nil {
		return domain.Location{}, fmt.Errorf("%w", err)
	}

	return savedLocation, nil
}

// GetAllLocation implement InLocationService interface
func (i *InLocationServiceImpl) GetAllLocation(ctx context.Context) ([]domain.Location, error) {
	locations, err := i.output.GetAllLocation(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return locations, nil
}
