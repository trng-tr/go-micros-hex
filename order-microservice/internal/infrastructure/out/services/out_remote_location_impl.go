package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/dtos"
)

type RemoteLocationServiceImpl struct {
	baseUrl string
}

func NewRemoteLocationServiceImpl(baseUrl string) *RemoteLocationServiceImpl {
	return &RemoteLocationServiceImpl{baseUrl: baseUrl}
}

func (o *RemoteLocationServiceImpl) GetRemoteLocationByID(ctx context.Context, locationID int64) (domain.Location, error) {
	remoteApiUrl := fmt.Sprintf(o.baseUrl+"/locations/%d", locationID)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, remoteApiUrl, nil)
	if err != nil {
		return domain.Location{}, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return domain.Location{}, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return domain.Location{}, errors.New("remote stock not found")
	}
	if response.StatusCode != http.StatusOK {
		return domain.Location{}, fmt.Errorf("remote stock service error: status %d", response.StatusCode)
	}
	// Decoder le remote dto
	var remoteLocationResponse dtos.LocationResponse
	if err := json.NewDecoder(response.Body).Decode(&remoteLocationResponse); err != nil {
		return domain.Location{}, err
	}

	return domain.Location{
		ID:          remoteLocationResponse.ID,
		Ville:       remoteLocationResponse.Ville,
		Description: remoteLocationResponse.Description,
	}, nil
}
