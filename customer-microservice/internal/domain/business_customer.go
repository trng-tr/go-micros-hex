package domain

import "time"

type BusinessCustomer struct {
	ID          int64
	Firstname   string
	Lastname    string
	Genda       Genda
	Email       string
	PhoneNumber string
	Status      Status
	AddressID   int64
	CreatedAt   time.Time //ex:2026-01-07 14:42:18.123456789 +0000 UTC
	UpdatedAt   string
}

type Genda string

const (
	Femal Genda = "F"
	Male  Genda = "M"
)

type Status string

const (
	Active    Status = "ACTIVE"
	Suspended Status = "SUSPENDED"
	Deleted   Status = "DELETED"
)

type BusinessAddress struct {
	ID           int64
	StreetNumber string
	StreetName   string
	ZipCode      string
	City         string
	Region       string
	Country      string
	Complement   string
}
