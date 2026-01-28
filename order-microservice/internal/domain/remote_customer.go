package domain

type Customer struct {
	ID          int64
	Firstname   string
	Lastname    string
	Genda       Genda
	Email       string
	PhoneNumber string
	Status      CustomerStatus
}

type Genda string

const (
	Female Genda = "F"
	Male   Genda = "M"
)

type CustomerStatus string

const Active CustomerStatus = "ACTIVE"
