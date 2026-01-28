package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trng-tr/customer-microservice/internal/application/ports/in"
	"github.com/trng-tr/customer-microservice/internal/domain"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/in/http/dtos"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/in/http/mappers"
)

/*
AddressHandlerServiceImpl implement interface AddressHandlerService
I inject Input port to access to usecase
*/
type AddressHandlerServiceImpl struct {
	inputPortSvc in.InAddressService
}

// NewAddressHandlerServiceImpl DI by constructor
func NewAddressHandlerServiceImpl(inputPortSvc in.InAddressService) *AddressHandlerServiceImpl {
	return &AddressHandlerServiceImpl{inputPortSvc: inputPortSvc}
}

// AddressHandleCreateService implement interface AddressHandlerService
func (ahsi AddressHandlerServiceImpl) AddressHandleCreate(c *gin.Context) {
	checkPostRequest(c)
	var addressRequest dtos.AddressRequest
	if err := c.ShouldBindJSON(&addressRequest); err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}
	var businessAddress domain.Address = mappers.ToBusinessAddress(addressRequest)
	createdAddress, err := ahsi.inputPortSvc.CreateAddress(c.Request.Context(), businessAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
	}

	var response dtos.AddressResponse = mappers.ToAddressResponse(createdAddress)

	c.JSON(http.StatusCreated, response)
}

// AddressHandleGetByIdService implement interface AddressHandlerService
func (ahsi *AddressHandlerServiceImpl) AddressHandleGetById(c *gin.Context) {
	checkGetMethod(c)
	addressId, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}

	businessAddress, err := ahsi.inputPortSvc.GetAddressByID(c, addressId)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}

	var response dtos.AddressResponse = mappers.ToAddressResponse(businessAddress)

	c.JSON(http.StatusOK, response)
}

// AddressHandleGetAll implement interface AddressHandlerService
func (h *AddressHandlerServiceImpl) AddressHandleGetAll(c *gin.Context) {
	checkGetMethod(c)
	addresses, err := h.inputPortSvc.GetAllAddresses(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	var addressesResponses []dtos.AddressResponse = make([]dtos.AddressResponse, 0, len(addresses))
	for _, a := range addresses {
		var r = mappers.ToAddressResponse(a)
		addressesResponses = append(addressesResponses, r)
	}
	c.JSON(http.StatusOK, addressesResponses)
}

// AddressHandleDelete implement interface AddressHandlerService
func (h *AddressHandlerServiceImpl) AddressHandleDelete(c *gin.Context) {
	checkDeleteMethod(c)
	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}

	if err := h.inputPortSvc.DeleteAddress(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}

	c.JSON(http.StatusOK, dtos.NewResponse(success, "successfuly deleted"))
}
