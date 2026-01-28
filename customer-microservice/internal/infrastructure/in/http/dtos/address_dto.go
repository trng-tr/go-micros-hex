package dtos

type AddressRequest struct {
	StreetNumber *string `json:"street_number,omitempty" binding:"omitempty"`
	StreetName   string  `json:"street_name" binding:"required"`
	ZipCode      string  `json:"zip_code" binding:"required"`
	City         string  `json:"city" binding:"required"`
	Region       string  `json:"region" binding:"required"`
	Country      string  `json:"country" binding:"required"`
	Complement   *string `json:"complement,omitempty"`
}

type AddressResponse struct {
	ID           int64   `json:"id"`
	StreetNumber *string `json:"street_number,omitempty"`
	StreetName   string  `json:"street_name"`
	ZipCode      string  `json:"zip_code"`
	City         string  `json:"city"`
	Region       string  `json:"region"`
	Country      string  `json:"country"`
	Complement   *string `json:"complement,omitempty"`
}
