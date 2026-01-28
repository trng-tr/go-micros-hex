package repositories

import (
	"context"
	"database/sql"

	"github.com/trng-tr/product-microservice/internal/infrastructure/out/models"
)

// LocationRepository implement interface Repository
type LocationRepository struct {
	db *sql.DB
}

func NewLocationRepository(db *sql.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

// Save implement interface Repository
func (l *LocationRepository) Save(ctx context.Context, o models.Location) (models.Location, error) {
	query := `INSERT INTO locations(ville,description,created_at,updated_at)
	VALUES ($1,$2,$3,$4)
	RETURNING id`
	if err := l.db.QueryRowContext(ctx, query, o.Ville, o.Description, o.CreatedAt,
		o.UpdatedAt).Scan(&o.ID); err != nil {
		return models.Location{}, err
	}

	return o, nil
}

// FindByID implement interface Repository
func (l *LocationRepository) FindByID(ctx context.Context, id int64) (models.Location, error) {
	query := `SELECT id,ville,description,created_at,updated_at 
	FROM locations
	WHERE id=$1`
	var model models.Location
	if err := l.db.QueryRowContext(ctx, query, id).Scan(
		&model.ID, &model.Ville, &model.Description, &model.CreatedAt, &model.UpdatedAt,
	); err != nil {
		return models.Location{}, err
	}

	return model, nil
}

// FindAll implement interface Repository
func (l *LocationRepository) FindAll(ctx context.Context) ([]models.Location, error) {
	query := `SELECT id,ville,description,created_at,updated_at FROM locations`
	rows, err := l.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []models.Location = make([]models.Location, 0)
	for rows.Next() {
		var model models.Location
		if err := rows.Scan(&model.ID, &model.Ville, &model.Description, &model.CreatedAt, &model.UpdatedAt); err != nil {
			return nil, err
		}
		data = append(data, model)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
