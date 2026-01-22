package services

import (
	"context"

	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/out/mappers"
	"github.com/trng-tr/product-microservice/internal/infrastructure/out/models"
)

// OutStockServiceImpl implement OutStockService interface
type OutStockServiceImpl struct {
	repo StockRepository //DI
}

// NewOutOutStockServiceImpl DI by constructor
func NewOutOutStockServiceImpl(repo StockRepository) *OutStockServiceImpl {
	return &OutStockServiceImpl{repo: repo}
}

// CreateStock implement OutStockService interface
func (o *OutStockServiceImpl) CreateStock(ctx context.Context, stk domain.Stock) (domain.Stock, error) {
	model, err := o.repo.SaveO(ctx, mappers.ToStockModel(stk))
	if err != nil {
		return domain.Stock{}, err
	}
	return utilMapp(model), nil
}

// GetStockByID implement OutStockService interface
func (o *OutStockServiceImpl) GetStockByID(ctx context.Context, id int64) (domain.Stock, error) {
	model, err := o.repo.FindOByID(ctx, id)
	if err != nil {
		return domain.Stock{}, err
	}
	return utilMapp(model), nil
}

// GetAllStocks implement OutStockService interface
func (o *OutStockServiceImpl) GetAllStocks(ctx context.Context) ([]domain.Stock, error) {
	models, err := o.repo.FindAllO(ctx)
	if err != nil {
		return nil, err
	}
	var bsStocks = make([]domain.Stock, 0, len(models))
	for _, model := range models {
		bsStocks = append(bsStocks, utilMapp(model))
	}

	return bsStocks, nil
}

// UpdateStockQuantity implement OutStockService interface
func (o *OutStockServiceImpl) UpdateStockQuantity(ctx context.Context, stock domain.Stock) (domain.Stock, error) {
	model, err := o.repo.UpdateStockQuantity(ctx, stock.ProductID, stock.Quantity)
	if err != nil {
		return domain.Stock{}, err
	}
	return utilMapp(model), nil
}

func utilMapp(model models.StockModel) domain.Stock {
	return mappers.ToBusinessStock(model)
}

func (o *OutStockServiceImpl) GetStockByProductID(ctx context.Context, productID int64) (domain.Stock, error) {
	model, err := o.repo.FindStockByProductID(ctx, productID)
	if err != nil {
		return domain.Stock{}, err
	}
	return utilMapp(model), nil
}
