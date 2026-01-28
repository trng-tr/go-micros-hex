package models

import (
	"database/sql"
	"time"
)

// CustomerModel model of data for table customers
type CustomerModel struct {
	ID          int64
	Firstname   string
	Lastname    string
	Genda       string
	Email       string
	PhoneNumber string
	Status      string
	AddressID   int64
	CreatedAt   time.Time    //ex:2026-01-07 14:42:18.123456789 +0000 UTC
	UpdatedAt   sql.NullTime //because it is not mandatory
}
