package usecase

import (
	"context"

	"github.com/trng-tr/order-microservice/internal/application/out"
	"github.com/trng-tr/order-microservice/internal/domain"
)

type RemoteLocationServiceImpl struct {
	outputPort out.RemoteLocationService
}

func NewRemoteLocationServiceImpl(outputPort out.RemoteLocationService) *RemoteLocationServiceImpl {
	return &RemoteLocationServiceImpl{outputPort: outputPort}
}

func (in *RemoteLocationServiceImpl) GetRemoteLocationByID(ctx context.Context, locationID int64) (domain.Location, error) {
	remoteLocation, err := in.outputPort.GetRemoteLocationByID(ctx, locationID)
	if err != nil {
		return domain.Location{}, err
	}

	return remoteLocation, nil
}
