package repositories

import (
	"context"
	"database/sql"

	"github.com/trng-tr/order-microservice/internal/infrastructure/out/models"
)

// OrderLineRepoImpl implements interface OrderLineRepo
type OrderLineRepoImpl struct {
	db *sql.DB
}

// NewOrderLineRepoImpl injection par constructeur
func NewOrderLineRepoImpl(db *sql.DB) *OrderLineRepoImpl {
	return &OrderLineRepoImpl{db: db}
}

// FindByID implement OrderLineRepo
func (o *OrderLineRepoImpl) FindByID(ctx context.Context, id int64) (models.OrderLineModel, error) {
	query := `SELECT id,order_id,product_id,quantity FROM orderlines WHERE id=$1`
	var model models.OrderLineModel
	if err := o.db.QueryRowContext(ctx, query, id).Scan(
		&model.ID, &model.OrderID, &model.ProductID, &model.Quantity,
	); err != nil {
		return models.OrderLineModel{}, err
	}

	return model, nil
}

// FindAll implement OrderLineRepo
func (o *OrderLineRepoImpl) FindAll(ctx context.Context) ([]models.OrderLineModel, error) {
	query := "SELECT id,order_id,product_id,quantity FROM orderlines"
	rows, err := o.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []models.OrderLineModel = make([]models.OrderLineModel, 0)
	for rows.Next() {
		var model models.OrderLineModel
		if err := rows.Scan(&model.ID, &model.OrderID, &model.ProductID, &model.Quantity); err != nil {
			return nil, err
		}
		data = append(data, model)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

// Delete implement OrderLineRepo
func (o *OrderLineRepoImpl) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM orderlines WHERE id=$1"
	results, err := o.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// Update implement OrderLineRepo
func (o *OrderLineRepoImpl) Update(ctx context.Context, id int64, quantity int64) (models.OrderLineModel, error) {
	query := `UPDATE orderlines
	SET quantity=$2
	WHERE id=$1
	RETURNING id,order_id,product_id,quantity`
	var model models.OrderLineModel
	if err := o.db.QueryRowContext(ctx, query, id, quantity).Scan(
		&model.ID, &model.OrderID, &model.ProductID, &model.Quantity); err != nil {
		return models.OrderLineModel{}, err
	}

	return model, nil
}

func (o *OrderLineRepoImpl) FindAllByOrderID(ctx context.Context, orderID int64) ([]models.OrderLineModel, error) {
	query := `SELECT id,order_id,product_id,quantity
	FROM orderlines
	WHERE order_id=$1`
	rows, err := o.db.QueryContext(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orderLines []models.OrderLineModel = make([]models.OrderLineModel, 0)
	for rows.Next() {
		var model models.OrderLineModel
		if err := rows.Scan(&model.ID, &model.OrderID, &model.ProductID, &model.Quantity); err != nil {
			return nil, err
		}
		orderLines = append(orderLines, model)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orderLines, nil
}
