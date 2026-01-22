package services

import (
	"bytes"
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
func (o *RemoteProductServiceImpl) GetRemoteProductByID(ctx context.Context, id int64) (domain.Product, error) {
	remoteApiUrl := fmt.Sprintf(o.baseUrl+"/products/%d", id)
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

// GetRemoteStockByProductID implement interface
func (o *RemoteProductServiceImpl) GetRemoteStockByProductID(ctx context.Context, prodID int64) (domain.Stock, error) {
	baseUrl := fmt.Sprintf(o.baseUrl+"/products/%d/stock", prodID)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, baseUrl, nil)
	if err != nil {
		return domain.Stock{}, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return domain.Stock{}, err
	}
	defer response.Body.Close()
	var stockResponse dtos.StockResponse
	if err := json.NewDecoder(response.Body).Decode(&stockResponse); err != nil {
		return domain.Stock{}, err
	}

	return utilMap(stockResponse), nil
}

// SetRemoteStockQuantity implement interface
func (o *RemoteProductServiceImpl) SetRemoteStockQuantity(ctx context.Context, productID int64, stock domain.Stock) error {
	baseUrl := fmt.Sprintf(o.baseUrl+"/stocks/set-qte/%d", productID)

	// equivalent of remote request 👇
	stockQuantityRequest := struct {
		Quantity int64 `json:"quantity"`
	}{Quantity: stock.Quantity}

	//encode in json onject 👇
	body, err := json.Marshal(stockQuantityRequest)
	if err != nil {
		return err
	}

	// create request qith context 👇
	request, err := http.NewRequestWithContext(ctx, http.MethodPut, baseUrl, bytes.NewReader(body))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	var stockResponse dtos.StockResponse
	if err := json.NewDecoder(resp.Body).Decode(&stockResponse); err != nil {
		return err
	}

	return nil
}

func utilMap(stockResponse dtos.StockResponse) domain.Stock {
	return domain.Stock{
		ID:        stockResponse.ID,
		Name:      stockResponse.Name,
		ProductID: stockResponse.ProductID,
		Quantity:  stockResponse.Quantity,
	}
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
