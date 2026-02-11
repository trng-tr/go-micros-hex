package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/trng-tr/product-microservice/internal/application/ports/out"
	"github.com/trng-tr/product-microservice/internal/domain"
)

// InProductServiceImpl implement InProductService
type InProductServiceImpl struct {
	outputPort out.OutProductService       //DI
	outputUuid out.OutUuidGeneratorService //DI
}

// NewInProductServiceImpl DI par cosntructor
func NewInProductServiceImpl(outputPort out.OutProductService, outputUuid out.OutUuidGeneratorService) *InProductServiceImpl {
	return &InProductServiceImpl{outputPort: outputPort, outputUuid: outputUuid}
}

// CreateProduct implement interface InProductService
func (i *InProductServiceImpl) SaveProduct(ctx context.Context, prod domain.Product) (domain.Product, error) {
	inputFieds := map[string]string{
		"category":    string(prod.Category),
		"prod_name":   prod.ProductName,
		"description": prod.Description,
	}

	if err := checkInputs1(inputFieds); err != nil {
		return domain.Product{}, err
	}

	if err := checkProdCategory(prod.Category); err != nil {
		return domain.Product{}, err
	}

	if err := checkPrice(prod.Price); err != nil {
		return domain.Product{}, err
	}
	prod.GenerateCreatedAt()
	//call OutUuidGeneratorService to generate uuid and set sku to product 👇
	var sku = i.outputUuid.GenerateUuid()
	prod.GenerateSku(prod.Category, prod.ProductName, sku)
	prod.IsActive = true
	// call output service to save product
	savdProd, err := i.outputPort.SaveProduct(ctx, prod)
	if err != nil {
		return domain.Product{}, fmt.Errorf("%w", err)
	}
	return savdProd, nil
}

func (i *InProductServiceImpl) GetProductByID(ctx context.Context, productID int64) (domain.Product, error) {
	if err := checkInputId(productID); err != nil {
		return domain.Product{}, err
	}
	savedPorduct, err := i.outputPort.GetProductByID(ctx, productID)
	if err != nil {
		return domain.Product{}, err
	}

	return savedPorduct, nil
}

// GetAllProducts implement interface InProductService
func (i *InProductServiceImpl) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	products, err := i.outputPort.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, fmt.Errorf("%w", errNoDataRegistered)
	}
	return products, nil
}

// PatchProduct implement interface InProductService
func (i *InProductServiceImpl) PatchProduct(ctx context.Context, productID int64, patch domain.PatchProduct) (domain.Product, error) {
	inputsvalues := map[string]int64{
		"product_id": productID,
	}
	if err := checkInputs2(inputsvalues); err != nil {
		return domain.Product{}, err
	}
	product, err := i.outputPort.GetProductByID(ctx, productID)
	if err != nil {
		return domain.Product{}, err
	}
	product.ApplyPatchMapper(patch)
	// call outputservice to save changes
	savedProd, err := i.outputPort.PatchProduct(ctx, productID, product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("%w:%v", errSavingObject, err)
	}

	return savedProd, nil
}

// DeleteProduct implement interface InProductService
func (i *InProductServiceImpl) DeleteProduct(ctx context.Context, productID int64) error {
	inputsvalues := map[string]int64{
		"product_id": productID,
	}
	if err := checkInputs2(inputsvalues); err != nil {
		return err
	}
	if _, err := i.outputPort.GetProductByID(ctx, productID); err != nil {
		return fmt.Errorf("%w:%v", errObjectNotFound, err)
	}

	if err := i.outputPort.DeleteProduct(ctx, productID); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// GetProductBySku implement interface InProductService
func (i *InProductServiceImpl) GetProductBySku(ctx context.Context, sku string) (domain.Product, error) {
	if sku == "" {
		return domain.Product{}, errors.New("error: sku is impty")
	}
	product, err := i.outputPort.GetProductBySku(ctx, sku)
	if err != nil {
		return domain.Product{}, fmt.Errorf("%w with sku %s", errObjectNotFound, sku)
	}

	return product, nil
}
