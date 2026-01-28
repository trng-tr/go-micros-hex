package services

import (
	"context"

	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/out/mappers"
)

// OutProductServiceImpl struct to implement OutProductService port
type OutProductServiceImpl struct {
	repo ProductRepository //DI
}

// NewOutProductServiceImpl DI by constructeur
func NewOutProductServiceImpl(repo ProductRepository) *OutProductServiceImpl {
	return &OutProductServiceImpl{repo: repo}
}

// SaveProduct implement output port OutProductService
func (o *OutProductServiceImpl) SaveProduct(ctx context.Context, prod domain.Product) (domain.Product, error) {
	var model = mappers.ToProductModel(prod)
	savedModel, err := o.repo.Save(ctx, model)
	if err != nil {
		return domain.Product{}, err
	}
	return mappers.ToBusinessProduct(savedModel), nil
}

// SaveProduct implement output port OutProductService
func (o *OutProductServiceImpl) GetProductByID(ctx context.Context, id int64) (domain.Product, error) {
	model, err := o.repo.FindByID(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return mappers.ToBusinessProduct(model), nil
}

// GetAllProducts implement output port OutProductService
func (o *OutProductServiceImpl) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
	models, err := o.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var bsProducts = make([]domain.Product, 0, len(models))
	for _, model := range models {
		bsProducts = append(bsProducts, mappers.ToBusinessProduct(model))
	}

	return bsProducts, nil
}

// PatchProduct implement output port OutProductService
func (o *OutProductServiceImpl) PatchProduct(ctx context.Context, id int64, product domain.Product) (domain.Product, error) {
	model, err := o.repo.PatchProduct(ctx, id, mappers.ToProductModel(product))
	if err != nil {

		return domain.Product{}, err
	}

	return mappers.ToBusinessProduct(model), nil
}

// DeleteProduct implement output port OutProductService
func (o *OutProductServiceImpl) DeleteProduct(ctx context.Context, id int64) error {
	if err := o.repo.DeleteProduct(ctx, id); err != nil {
		return err
	}
	return nil
}

// GetProductBySku implement output port OutProductService
func (o *OutProductServiceImpl) GetProductBySku(ctx context.Context, sku string) (domain.Product, error) {
	model, err := o.repo.FindProductBySku(ctx, sku)
	if err != nil {
		return domain.Product{}, err
	}

	return mappers.ToBusinessProduct(model), nil
}
