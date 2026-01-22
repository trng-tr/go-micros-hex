package repositories

import (
	"context"
	"database/sql"

	"github.com/trng-tr/product-microservice/internal/infrastructure/out/models"
)

// StockRepositoryImpl wants to implement interface ProductRepository
type StockRepositoryImpl struct {
	db *sql.DB
}

// NewStockRepositoryImpl DI by constructor
func NewStockRepositoryImpl(db *sql.DB) *StockRepositoryImpl {
	return &StockRepositoryImpl{db: db}
}

// SaveO implement interface StockRepository
func (s *StockRepositoryImpl) SaveO(ctx context.Context, o models.StockModel) (models.StockModel, error) {
	var query = `INSERT INTO stocks (name,product_id,quantity,updated_at)
	VALUES($1,$2,$3,$4)
	RETURNING id`
	if err := s.db.QueryRowContext(
		ctx,
		query,
		o.Name,
		o.ProductID,
		o.Quantity,
		o.UpdatedAt,
	).Scan(&o.ID); err != nil {
		return models.StockModel{}, err
	}

	return o, nil
}

// FindOByID implement interface StockRepository
func (s *StockRepositoryImpl) FindOByID(ctx context.Context, id int64) (models.StockModel, error) {
	query := `SELECT id,name,product_id,quantity,updated_at
	FROM stocks
	WHERE id=$1`
	var stock models.StockModel
	if err := s.db.QueryRowContext(ctx, query, id).Scan(
		&stock.ID,
		&stock.Name,
		&stock.ProductID,
		&stock.Quantity,
		&stock.UpdatedAt,
	); err != nil {
		return models.StockModel{}, err
	}

	return stock, nil
}

// FindAllO implement interface StockRepository
func (s *StockRepositoryImpl) FindAllO(ctx context.Context) ([]models.StockModel, error) {
	var query string = "SELECT id,name,product_id,quantity,updated_at FROM stocks"
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var stocks []models.StockModel = make([]models.StockModel, 0)
	for rows.Next() {
		var stock models.StockModel
		if err := rows.Scan(
			&stock.ID,
			&stock.Name,
			&stock.ProductID,
			&stock.Quantity,
			&stock.UpdatedAt,
		); err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stocks, nil
}

// DeleteO implement interface StockRepository
func (s *StockRepositoryImpl) DeleteO(ctx context.Context, id int64) error {
	query := "DELETE FROM stocks WHERE id=$1"
	result, err := s.db.ExecContext(ctx, query, id)
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

// UpdateStockQuantity implement interface StockRepository, set quantity
func (s *StockRepositoryImpl) UpdateStockQuantity(ctx context.Context, productID int64, quantity int64) (models.StockModel, error) {
	var query = `UPDATE stocks 
	SET quantity = $2
	WHERE product_id = $1
	RETURNING id,name,product_id,quantity,updated_at`
	var newStock models.StockModel
	if err := s.db.QueryRowContext(ctx, query, productID, quantity).Scan(
		&newStock.ID, &newStock.Name, &newStock.ProductID, &newStock.Quantity, &newStock.UpdatedAt); err != nil {
		return models.StockModel{}, err
	}

	return newStock, nil
}

func (s *StockRepositoryImpl) FindStockByProductID(ctx context.Context, productID int64) (models.StockModel, error) {
	query := `SELECT id,name,product_id,quantity,updated_at
	FROM stocks
	WHERE product_id=$1`
	var stock models.StockModel
	if err := s.db.QueryRowContext(ctx, query, productID).Scan(
		&stock.ID,
		&stock.Name,
		&stock.ProductID,
		&stock.Quantity,
		&stock.UpdatedAt,
	); err != nil {
		return models.StockModel{}, err
	}

	return stock, nil
}
