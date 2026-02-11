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

// RemoteProductServiceImpl Implemnet intreface RemoteProductService
type RemoteProductServiceImpl struct {
	baseUrl string
}

// NewRemoteProductServiceImpl DI by constructor
func NewRemoteProductServiceImpl(url string) *RemoteProductServiceImpl {
	return &RemoteProductServiceImpl{baseUrl: url}
}

// GetRemoteProductByID implement interface
func (o *RemoteProductServiceImpl) GetRemoteProductByID(ctx context.Context, productID int64) (domain.Product, error) {
	remoteApiUrl := fmt.Sprintf(o.baseUrl+"/products/%d", productID)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, remoteApiUrl, nil)
	if err != nil {
		return domain.Product{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return domain.Product{}, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return domain.Product{}, errors.New("remote product not found")
	}
	if response.StatusCode != http.StatusOK {
		return domain.Product{}, fmt.Errorf("remote product service error: status %d", response.StatusCode)
	}
	// Decoder le remote dto
	var remoteProductResponse dtos.ProductResponse
	if err := json.NewDecoder(response.Body).Decode(&remoteProductResponse); err != nil {
		return domain.Product{}, err
	}

	domainProduct := toDomainProduct(remoteProductResponse)

	return domainProduct, nil
}

func toDomainProduct(remoteResponse dtos.ProductResponse) domain.Product {
	return domain.Product{
		ID:          remoteResponse.ID,
		Sku:         remoteResponse.Sku,
		Category:    domain.Category(remoteResponse.Category),
		ProductName: remoteResponse.ProductName,
		Description: remoteResponse.Description,
		Price: domain.Price{
			UnitPrice: remoteResponse.PriceResponse.UnitPrice,
			Currency:  domain.Currency(remoteResponse.PriceResponse.Currency),
		},
		IsActive: remoteResponse.IsActive,
	}
}
