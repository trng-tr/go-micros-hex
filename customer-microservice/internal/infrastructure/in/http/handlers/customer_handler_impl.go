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
CustomerHandlerServiceImpl implement interface CustomerHandlerService
I inject Input port to access to usecase
*/
type CustomerHandlerServiceImpl struct {
	cInputPortSvc in.InCustomerService
	aInputPortSvc in.InAddressService
}

// NewCustomerHandlerServiceImpl DI by constructor
func NewCustomerHandlerServiceImpl(cInputPortSvc in.InCustomerService, aInputPortSvc in.InAddressService) *CustomerHandlerServiceImpl {
	return &CustomerHandlerServiceImpl{cInputPortSvc: cInputPortSvc, aInputPortSvc: aInputPortSvc}
}

// CustomerHandleCreateService implement CustomerHandlerService
func (ch *CustomerHandlerServiceImpl) CustomerHandleCreate(c *gin.Context) {
	checkPostRequest(c)
	var request dtos.CustomerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}
	bsCustomer, err := ch.cInputPortSvc.CreateCustomer(c.Request.Context(), mappers.ToBusinessCustomer(request))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	bsAddress, err := ch.aInputPortSvc.GetAddressByID(c.Request.Context(), bsCustomer.AddressID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, buildCustomerResponse(bsCustomer, bsAddress))
}

// CustomerHandleGetById implement CustomerHandlerService
func (ch *CustomerHandlerServiceImpl) CustomerHandleGetById(c *gin.Context) {
	checkGetMethod(c)
	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}
	bsCustomer, err := ch.cInputPortSvc.GetCustomerByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	bsAddress, err := ch.aInputPortSvc.GetAddressByID(c.Request.Context(), bsCustomer.AddressID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}

	c.JSON(http.StatusOK, buildCustomerResponse(bsCustomer, bsAddress))

}

// CustomerHandleGetAll implement CustomerHandlerService
func (ch *CustomerHandlerServiceImpl) CustomerHandleGetAll(c *gin.Context) {
	checkGetMethod(c)
	bsCustomers, err := ch.cInputPortSvc.GetAllCustomers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	var responses []dtos.CustomerResponse = make([]dtos.CustomerResponse, 0, len(bsCustomers))
	for _, bsCustomer := range bsCustomers {
		bsAddress, err := ch.aInputPortSvc.GetAddressByID(c.Request.Context(), bsCustomer.AddressID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
			return
		}
		responses = append(responses, buildCustomerResponse(bsCustomer, bsAddress))
	}

	c.JSON(http.StatusOK, responses)
}

// CustomerHandlePatch implement CustomerHandlerService
func (ch *CustomerHandlerServiceImpl) CustomerHandlePatch(c *gin.Context) {
	if c.Request.Method != http.MethodPatch {
		c.JSON(http.StatusMethodNotAllowed, dtos.NewResponse(fail, errMethodNotAllowed.Error()))
		return
	}
	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}

	var patchRequest dtos.CustomerPatchRequest
	if err := c.ShouldBindJSON(&patchRequest); err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}
	var bsPatch domain.PatchBusinessCustomer = mappers.ToDomainPatch(patchRequest)

	bsCustomer, err := ch.cInputPortSvc.PatchCustomer(c.Request.Context(), id, bsPatch)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}
	bsAddress, err := ch.aInputPortSvc.GetAddressByID(c.Request.Context(), bsCustomer.AddressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}
	c.JSON(http.StatusOK, buildCustomerResponse(bsCustomer, bsAddress))
}

// CustomerHandleDelete implement CustomerHandlerService
func (ch *CustomerHandlerServiceImpl) CustomerHandleDelete(c *gin.Context) {
	checkDeleteMethod(c)

	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}
	if err := ch.cInputPortSvc.DeleteCustomer(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	c.JSON(http.StatusOK, dtos.NewResponse(success, "successfuly deleted"))
}

// buildCustomerResponse util function
func buildCustomerResponse(bsC domain.Customer, bsA domain.Address) dtos.CustomerResponse {
	return mappers.ToCustomerResponse(bsC, bsA)
}
