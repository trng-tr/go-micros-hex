package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/trng-tr/product-microservice/internal/application/ports/out"
	"github.com/trng-tr/product-microservice/internal/domain"
)

// InStockServiceImpl implement InStockService
type InStockServiceImpl struct {
	stockOutService out.OutStockService
	prodOutService  out.OutProductService
}

// NewInStockServiceImpl DI par constructeur
func NewInStockServiceImpl(out1 out.OutStockService, out2 out.OutProductService) *InStockServiceImpl {
	return &InStockServiceImpl{stockOutService: out1, prodOutService: out2}
}

// CreateStock implement interface InStockService
func (i *InStockServiceImpl) CreateStock(ctx context.Context, stk domain.Stock) (domain.Stock, error) {
	var inputFields = map[string]int64{
		"product_id": stk.ProductID,
		"quantity":   stk.Quantity,
	}
	if err := checkStockInputs(inputFields); err != nil {
		return domain.Stock{}, err
	}

	if err := checkStockName(stk.Name); err != nil {
		return domain.Stock{}, err
	}
	if _, err := i.prodOutService.GetProductByID(ctx, stk.ProductID); err != nil {
		return domain.Stock{}, fmt.Errorf("%w:%v", errObjectNotFound, err)
	}

	stk.UpdatedAt = time.Now()
	//call output service to register stock
	savedStock, err := i.stockOutService.CreateStock(ctx, stk)
	if err != nil {
		return domain.Stock{}, fmt.Errorf("%w,%v", errSavingObject, err)
	}

	return savedStock, nil
}

// GetStockByID implement interface InStockService
func (i *InStockServiceImpl) GetStockByID(ctx context.Context, id int64) (domain.Stock, error) {
	if err := checkInputId(id); err != nil {
		return domain.Stock{}, err
	}
	stock, err := i.stockOutService.GetStockByID(ctx, id)
	if err != nil {
		return domain.Stock{}, fmt.Errorf("%w:%v", errObjectNotFound, err)
	}

	return stock, nil
}

// GetAllStocks implement interface InStockService
func (i *InStockServiceImpl) GetAllStocks(ctx context.Context) ([]domain.Stock, error) {
	stocks, err := i.stockOutService.GetAllStocks(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w:%v", errPrintingObjects, err)
	}
	if len(stocks) == 0 {
		return nil, fmt.Errorf("%w", errNoDataRegistered)
	}

	return stocks, nil
}

// SetStockProductQuantity implement interface InStockService
func (i *InStockServiceImpl) SetStockQuantity(ctx context.Context, productID int64, newQuantity int64) (domain.Stock, error) {
	if err := checkInputId(productID); err != nil {
		return domain.Stock{}, err
	}
	if err := checkInputStockQty(newQuantity); err != nil {
		return domain.Stock{}, err
	}
	stock, err := i.GetStockByProductID(ctx, productID)
	if err != nil {
		return domain.Stock{}, err
	}
	stock.Quantity = newQuantity // replace old quantity
	//call output service to register new stock with new quantity
	savedStock, err := i.stockOutService.UpdateStockQuantity(ctx, stock)
	if err != nil {
		return domain.Stock{}, fmt.Errorf("%w:%v", errSavingObject, err)
	}
	return savedStock, nil
}

// IncreaseStockProductQuantity implement interface InStockService
func (i *InStockServiceImpl) IncreaseStockQuantity(ctx context.Context, productID int64, quantity int64) (domain.Stock, error) {
	if err := checkInputId(productID); err != nil {
		return domain.Stock{}, err
	}
	if err := checkInputStockQty(quantity); err != nil {
		return domain.Stock{}, err
	}
	stock, err := i.GetStockByProductID(ctx, productID)
	if err != nil {
		return domain.Stock{}, err
	}
	stock.Quantity += quantity //increase stock product quantity
	//call output service to register stock with updated quantity
	savedStock, err := i.stockOutService.UpdateStockQuantity(ctx, stock)

	if err != nil {
		return domain.Stock{}, fmt.Errorf("%w:%v", errSavingObject, err)
	}

	return savedStock, nil
}

// IncreaseStockProductQuantity implement interface InStockService
func (i *InStockServiceImpl) DecreaseStockQuantity(ctx context.Context, productID int64, quantity int64) (domain.Stock, error) {
	if err := checkInputId(productID); err != nil {
		return domain.Stock{}, err
	}
	if err := checkInputStockQty(quantity); err != nil {
		return domain.Stock{}, err
	}
	stock, err := i.GetStockByProductID(ctx, productID)
	if err != nil {
		return domain.Stock{}, err
	}
	// ne pas autorisé le stock negatif d'un produit
	if stock.Quantity-quantity < 0 {
		return domain.Stock{}, fmt.Errorf("%w", errInsufficientStock)
	}
	stock.Quantity -= quantity //increase stock product quantity
	//call output service to register stock with updated quantity
	savedStock, err := i.stockOutService.UpdateStockQuantity(ctx, stock)

	if err != nil {
		return domain.Stock{}, fmt.Errorf("%w:%v", errSavingObject, err)
	}

	return savedStock, nil
}

func (i *InStockServiceImpl) GetStockByProductID(ctx context.Context, productID int64) (domain.Stock, error) {
	if err := checkInputId(productID); err != nil {
		return domain.Stock{}, err
	}
	stock, err := i.stockOutService.GetStockByProductID(ctx, productID)
	if err != nil {
		return domain.Stock{}, fmt.Errorf("%w:%v", errOccured, err)
	}

	return stock, nil
}
