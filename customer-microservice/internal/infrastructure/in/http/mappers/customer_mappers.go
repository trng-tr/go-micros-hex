package mappers

import (
	"github.com/trng-tr/customer-microservice/internal/domain"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/in/http/dtos"
)

const dateFormat string = "2006-01-02 15:04:05"

func ToBusinessCustomer(request dtos.CustomerRequest) domain.Customer {
	return domain.Customer{
		Firstname:   request.Firstname,
		Lastname:    request.Lastname,
		Genda:       domain.Genda(request.Genda),
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		AddressID:   request.AddressID,
	}
}

func ToCustomerResponse(bs domain.Customer, bsAddress domain.Address) dtos.CustomerResponse {
	var updatedDate *string
	if bs.UpdatedAt != nil {
		s := bs.UpdatedAt.Format(dateFormat)
		updatedDate = &s
	}
	return dtos.CustomerResponse{
		ID:                   bs.ID,
		Firstname:            bs.Firstname,
		Lastname:             bs.Lastname,
		Genda:                string(bs.Genda),
		Email:                bs.Email,
		PhoneNumber:          bs.PhoneNumber,
		Status:               string(bs.Status),
		LightAddressResponse: ToLightAddressResponse(bsAddress),
		CreatedAt:            bs.CreatedAt.Format(dateFormat),
		UpdatedAt:            updatedDate,
	}
}

func ToDomainPatch(patchRequest dtos.CustomerPatchRequest) domain.PatchBusinessCustomer {
	return domain.PatchBusinessCustomer{
		Firstname:   patchRequest.Firstname,
		Lastname:    patchRequest.Lastname,
		Email:       patchRequest.Email,
		PhoneNumber: patchRequest.PhoneNumber,
		AddressID:   patchRequest.AddressID,
	}
}
