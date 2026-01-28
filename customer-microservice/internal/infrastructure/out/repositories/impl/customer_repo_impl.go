package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/trng-tr/customer-microservice/internal/infrastructure/out/models"
)

/*
CustomerRepositoryImpl implement CustomerRepository,
it uses by DI sql.DB
*/
type CustomerRepositoryImpl struct {
	db *sql.DB
}

// NewCustomerRepositoryImpl function constructor
func NewCustomerRepositoryImpl(db *sql.DB) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{db: db}
}

func (cri *CustomerRepositoryImpl) SaveO(ctx context.Context, o models.CustomerModel) (models.CustomerModel, error) {
	var query string = `
	INSERT INTO customers (firstname,lastname,genda,email,phone_number,status,address_id,created_at,updated_at)
	VALUES(
	$1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id`
	var err error = cri.db.QueryRowContext(
		ctx,
		query,
		o.Firstname,
		o.Lastname,
		o.Genda,
		o.Email,
		o.PhoneNumber,
		o.Status,
		o.AddressID,
		o.CreatedAt,
		o.UpdatedAt,
	).Scan(&o.ID)

	if err != nil {
		return models.CustomerModel{}, err
	}
	return o, nil
}

// FindOByID implement interface CustomerRepository
func (cri *CustomerRepositoryImpl) FindOByID(ctx context.Context, id int64) (models.CustomerModel, error) {
	var query string = `
	SELECT id,firstname,lastname,genda,email,phone_number,status,address_id,created_at,updated_at
	FROM customers
	WHERE id=$1`
	var customer models.CustomerModel
	var err error = cri.db.QueryRowContext(ctx, query, id).Scan(
		&customer.ID,
		&customer.Firstname,
		&customer.Lastname,
		&customer.Genda,
		&customer.Email,
		&customer.PhoneNumber,
		&customer.Status,
		&customer.AddressID,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)

	if err != nil {
		return models.CustomerModel{}, err
	}

	return customer, nil
}

// FindAllO implement interface CustomerRepository
func (cri *CustomerRepositoryImpl) FindAllO(ctx context.Context) ([]models.CustomerModel, error) {
	var query string = `
	SELECT id,firstname,lastname,genda,email,phone_number,status,address_id,created_at,updated_at
	 FROM customers`
	rows, err := cri.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var customers []models.CustomerModel = make([]models.CustomerModel, 0)
	for rows.Next() {
		var customer models.CustomerModel
		err := rows.Scan(
			&customer.ID,
			&customer.Firstname,
			&customer.Lastname,
			&customer.Genda,
			&customer.Email,
			&customer.Status,
			&customer.PhoneNumber,
			&customer.AddressID,
			&customer.CreatedAt,
			&customer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	// 👇 vérifie les erreurs survenues pendant l’itération
	if err := rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()

	return customers, nil
}

// UpdateO implement interface CustomerRepository
func (cri *CustomerRepositoryImpl) UpdateO(ctx context.Context, id int64, o models.CustomerModel) (models.CustomerModel, error) {
	query := `UPDATE customers
	SET
		firstname=$1,
		lastname=$2,
		email=$3,
		phone_number=$4,
		address_id=$5,
		updated_at=$6
	WHERE id=$7
	RETURNING id, firstname, lastname, genda, email, phone_number, status, address_id, created_at, updated_at`

	var updatedCustomer models.CustomerModel
	var err error = cri.db.QueryRowContext(
		ctx,
		query,
		o.Firstname,
		o.Lastname,
		o.Email,
		o.PhoneNumber,
		o.AddressID,
		o.UpdatedAt,
		id,
	).Scan(
		&updatedCustomer.ID,
		&updatedCustomer.Firstname,
		&updatedCustomer.Lastname,
		&updatedCustomer.Genda,
		&updatedCustomer.Email,
		&updatedCustomer.PhoneNumber,
		&updatedCustomer.Status,
		&updatedCustomer.AddressID,
		&updatedCustomer.CreatedAt,
		&updatedCustomer.UpdatedAt,
	)
	if err != nil {
		return models.CustomerModel{}, fmt.Errorf("an error occurred %w", err)
	}

	return updatedCustomer, nil
}

// DeleteO implement interface CustomerRepository
func (cri *CustomerRepositoryImpl) DeleteO(ctx context.Context, id int64) error {
	var query string = `DELETE FROM customers WHERE id=$1`
	result, err := cri.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("an error occurred %w", err)
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("an error occurred %w", err)
	}
	if rowAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
