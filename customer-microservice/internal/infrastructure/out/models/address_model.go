package models

import "database/sql"

//AddressModel model of data for table addresses
type AddressModel struct {
	ID           int64
	StreetNumber sql.NullString //because it is not mandatory
	StreetName   string
	ZipCode      string
	City         string
	Region       string
	Country      string
	Complement   sql.NullString //because it is not mandatory
}
