package domain

type Address struct {
	ID           int64
	StreetNumber *string
	StreetName   string
	ZipCode      string
	City         string
	Region       string
	Country      string
	Complement   *string //because it is not mandatory
}
