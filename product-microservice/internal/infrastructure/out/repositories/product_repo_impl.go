package repositories

import (
	"context"
	"database/sql"

	"github.com/trng-tr/product-microservice/internal/infrastructure/out/models"
)

// ProductRepositoryImpl wants to implement interface ProductRepository
type ProductRepositoryImpl struct {
	db *sql.DB //DI
}

// NewProductRepositoryImpl DI by constructor
func NewProductRepositoryImpl(db *sql.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

// SaveO implement interface ProductRepository
func (p *ProductRepositoryImpl) Save(ctx context.Context, o models.ProductModel) (models.ProductModel, error) {
	var query = `INSERT INTO products (sku,category,product_name,description,unit_price,currency,created_at,updated_at,is_active)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id`
	if err := p.db.QueryRowContext(
		ctx,
		query,
		o.Sku,
		o.Categoy,
		o.ProductName,
		o.Description,
		o.UnitPrice,
		o.Currency,
		o.CreatedAt,
		o.UpdatedAt,
		o.IsActive,
	).Scan(&o.ID); err != nil {
		return models.ProductModel{}, err
	}

	return o, nil
}

// FindOByID implement interface ProductRepository
func (p *ProductRepositoryImpl) FindByID(ctx context.Context, id int64) (models.ProductModel, error) {
	query := `SELECT id,sku,category,product_name,description,unit_price,currency,created_at,updated_at,is_active
	FROM products
	WHERE id=$1`
	var product models.ProductModel
	if err := p.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Sku,
		&product.Categoy,
		&product.ProductName,
		&product.Description,
		&product.UnitPrice,
		&product.Currency,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.IsActive,
	); err != nil {
		return models.ProductModel{}, err
	}

	return product, nil
}

// FindAllO implement interface ProductRepository
func (p *ProductRepositoryImpl) FindAll(ctx context.Context) ([]models.ProductModel, error) {
	var query string = `
	SELECT id,sku,category,product_name,description,unit_price,currency,created_at,updated_at,is_active
	 FROM products`
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var products []models.ProductModel = make([]models.ProductModel, 0)
	for rows.Next() {
		var product models.ProductModel
		if err := rows.Scan(
			&product.ID,
			&product.Sku,
			&product.Categoy,
			&product.ProductName,
			&product.Description,
			&product.UnitPrice,
			&product.Currency,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.IsActive,
		); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// DeleteO implement interface ProductRepository
func (p *ProductRepositoryImpl) DeleteProduct(ctx context.Context, id int64) error {
	query := "DELETE FROM products WHERE id=$1"
	result, err := p.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// PatchO implement interface ProductRepository
func (p *ProductRepositoryImpl) PatchProduct(ctx context.Context, id int64, o models.ProductModel) (models.ProductModel, error) {
	query := `UPDATE products
	SET
		product_name=$1,
		description=$2,
		unit_price=$3,
		updated_at=$4,
		is_active=$5
	WHERE id=$6
	RETURNING id,sku,category,product_name,description,unit_price,currency,created_at,updated_at,is_active`
	var product models.ProductModel
	if err := p.db.QueryRowContext(
		ctx,
		query,
		o.ProductName,
		o.Description,
		o.UnitPrice,
		o.UpdatedAt,
		o.IsActive,
		id,
	).Scan(
		&product.ID,
		&product.Sku,
		&product.Categoy,
		&product.ProductName,
		&product.Description,
		&product.UnitPrice,
		&product.Currency,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.IsActive,
	); err != nil {
		return models.ProductModel{}, err
	}

	return product, nil
}

// FindProductBySku implement interface ProductRepository
func (p *ProductRepositoryImpl) FindProductBySku(ctx context.Context, sku string) (models.ProductModel, error) {
	var query string = `
	SELECT id,sku,category,product_name,description,unit_price,currency,created_at,updated_at,is_active
	 FROM products
	 WHERE sku=$1`
	var prod models.ProductModel
	if err := p.db.QueryRowContext(ctx, query, sku).Scan(
		&prod.ID,
		&prod.Sku,
		&prod.Categoy,
		&prod.ProductName,
		&prod.Description,
		&prod.UnitPrice,
		&prod.Currency,
		&prod.CreatedAt,
		&prod.UpdatedAt,
		&prod.IsActive,
	); err != nil {
		return models.ProductModel{}, err
	}

	return prod, nil
}
