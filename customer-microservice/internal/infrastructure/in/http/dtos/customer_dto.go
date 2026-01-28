package dtos

// CustomerRequest create request
type CustomerRequest struct {
	Firstname   string `json:"firstname" binding:"required"`
	Lastname    string `json:"lastname" binding:"required"`
	Genda       string `json:"genda" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	AddressID   int64  `json:"address_id" binding:"required"`
}

// CustomerPatchRequest patch request
type CustomerPatchRequest struct {
	Firstname   *string `json:"firstname,omitempty" binding:"omitempty,min=2"`
	Lastname    *string `json:"lastname,omitempty" binding:"omitempty,min=2"`
	Email       *string `json:"email,omitempty" binding:"omitempty,min=5"`
	PhoneNumber *string `json:"phone_number,omitempty" binding:"omitempty,min=8,max=20"`
	AddressID   *int64  `json:"address_id,omitempty" binding:"omitempty,gt=0"`
}

type CustomerResponse struct {
	ID                   int64                `json:"id"`
	Firstname            string               `json:"firstname"`
	Lastname             string               `json:"lastname"`
	Genda                string               `json:"genda"`
	Email                string               `json:"email"`
	PhoneNumber          string               `json:"phone_number"`
	Status               string               `json:"status"`
	LightAddressResponse LightAddressResponse `json:"address"`
	CreatedAt            string               `json:"created_date"`
	UpdatedAt            *string              `json:"updated_date,omitempty"`
}

type LightAddressResponse struct {
	StreetNumber *string `json:"street_number,omitempty"`
	StreetName   string  `json:"street_name"`
	ZipCode      string  `json:"zip_code"`
	City         string  `json:"city"`
	Region       string  `json:"region"`
	Country      string  `json:"country"`
	Complement   *string `json:"complement,omitempty"`
}
