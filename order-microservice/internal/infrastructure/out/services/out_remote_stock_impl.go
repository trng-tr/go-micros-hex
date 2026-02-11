package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/dtos"
)

// RemoteStockServiceImpl implemnt interface RemoteStockService
type RemoteStockServiceImpl struct {
	baseUrl string
}

// NewRemoteStockServiceImpl injection par constructeur
func NewRemoteStockServiceImpl(baseUrl string) *RemoteStockServiceImpl {
	return &RemoteStockServiceImpl{baseUrl: baseUrl}
}

// GetRemoteStockByProductID implement interface
func (o *RemoteStockServiceImpl) GetRemoteStockByLocationIDAndProductID(ctx context.Context, locationID, prodID int64) (domain.Stock, error) {
	baseUrl := fmt.Sprintf(o.baseUrl+"/stocks/locations/%d/products/%d", locationID, prodID)
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
func (o *RemoteStockServiceImpl) SetRemoteStockQuantity(ctx context.Context, locationID, productID int64, stock domain.Stock) error {
	baseUrl := fmt.Sprintf(o.baseUrl+"/stocks/locations/%d/products/%d/set-qte", locationID, productID)

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
		ID:         stockResponse.ID,
		Name:       stockResponse.Name,
		ProductID:  stockResponse.ProductResponse.ID,
		LocationID: stockResponse.LocationResponse.ID,
		Quantity:   stockResponse.Quantity,
	}
}
