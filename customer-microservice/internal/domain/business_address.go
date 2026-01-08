package domain

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
