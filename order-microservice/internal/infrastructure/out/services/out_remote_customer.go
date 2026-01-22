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

// RemoteCustomerServiceImpl implement RemoteCustomerService
type RemoteCustomerServiceImpl struct {
	baseUrl string
}

// NewRemoteCustomerServiceImpl func construteur
func NewRemoteCustomerServiceImpl(baseUrl string) *RemoteCustomerServiceImpl {
	return &RemoteCustomerServiceImpl{baseUrl: baseUrl}
}

// GetRemoteOByID implements RemoteCustomerService
func (o *RemoteCustomerServiceImpl) GetRemoteCustomerByID(ctx context.Context, id int64) (domain.Customer, error) {
	remoteApiUrl := fmt.Sprintf(o.baseUrl+"/customers/%d", id)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, remoteApiUrl, nil)
	if err != nil {
		return domain.Customer{}, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return domain.Customer{}, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return domain.Customer{}, errors.New("remote customer not found")
	}
	if response.StatusCode != http.StatusOK {
		return domain.Customer{}, fmt.Errorf("remote customer service error: status %d", response.StatusCode)
	}
	// Decoder le remote dto
	var remoteCustomerResponse dtos.CustomerResponse
	if err := json.NewDecoder(response.Body).Decode(&remoteCustomerResponse); err != nil {
		return domain.Customer{}, err
	}

	domainCustomer := toDomainCustomer(remoteCustomerResponse)

	return domainCustomer, nil
}

func toDomainCustomer(dtoResp dtos.CustomerResponse) domain.Customer {
	return domain.Customer{
		ID:          dtoResp.ID,
		Firstname:   dtoResp.Firstname,
		Lastname:    dtoResp.Lastname,
		Genda:       domain.Genda(dtoResp.Genda),
		Email:       dtoResp.Email,
		PhoneNumber: dtoResp.PhoneNumber,
		Status:      domain.CustomerStatus(dtoResp.Status),
	}
}
