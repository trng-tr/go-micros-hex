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
func (s *StockRepositoryImpl) Save(ctx context.Context, o models.StockModel) (models.StockModel, error) {
	var query = `INSERT INTO stocks (name,product_id,location_id, quantity,updated_at)
	VALUES($1,$2,$3,$4,$5)
	RETURNING id`
	if err := s.db.QueryRowContext(
		ctx,
		query,
		o.Name,
		o.ProductID,
		o.LocationID,
		o.Quantity,
		o.UpdatedAt,
	).Scan(&o.ID); err != nil {
		return models.StockModel{}, err
	}

	return o, nil
}

// FindOByID implement interface StockRepository
func (s *StockRepositoryImpl) FindByID(ctx context.Context, id int64) (models.StockModel, error) {
	query := `SELECT id,name,product_id,location_id,quantity,updated_at
	FROM stocks
	WHERE id=$1`
	var stock models.StockModel
	if err := s.db.QueryRowContext(ctx, query, id).Scan(
		&stock.ID,
		&stock.Name,
		&stock.ProductID,
		&stock.LocationID,
		&stock.Quantity,
		&stock.UpdatedAt,
	); err != nil {
		return models.StockModel{}, err
	}

	return stock, nil
}

// FindAllO implement interface StockRepository
func (s *StockRepositoryImpl) FindAll(ctx context.Context) ([]models.StockModel, error) {
	var query string = "SELECT id,name,product_id,location_id,quantity,updated_at FROM stocks"
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
			&stock.LocationID,
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

// UpdateStockQuantity implement interface StockRepository, set quantity
func (s *StockRepositoryImpl) UpdateStockQuantity(ctx context.Context, productID, LocationID int64, quantity int64) (models.StockModel, error) {
	var query = `UPDATE stocks 
	SET quantity = $3
	WHERE product_id=$1 AND location_id = $2 
	RETURNING id,name,product_id,location_id,quantity,updated_at`
	var newStock models.StockModel
	if err := s.db.QueryRowContext(ctx, query, productID, LocationID, quantity).Scan(
		&newStock.ID, &newStock.Name, &newStock.ProductID, &newStock.LocationID, &newStock.Quantity, &newStock.UpdatedAt); err != nil {
		return models.StockModel{}, err
	}

	return newStock, nil
}

func (s *StockRepositoryImpl) FindStockByLocationIDAndProductID(ctx context.Context, locationID, productID int64) (models.StockModel, error) {
	query := `SELECT id,name,product_id,location_id,quantity,updated_at
	FROM stocks
	WHERE location_id=$1 AND product_id=$2`
	var stock models.StockModel
	if err := s.db.QueryRowContext(ctx, query, locationID, productID).Scan(
		&stock.ID,
		&stock.Name,
		&stock.ProductID,
		&stock.LocationID,
		&stock.Quantity,
		&stock.UpdatedAt,
	); err != nil {
		return models.StockModel{}, err
	}

	return stock, nil
}
