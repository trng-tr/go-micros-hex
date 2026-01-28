package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trng-tr/product-microservice/internal/application/ports/in"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/dtos"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/mappers"
)

// LocationHandlerServiceImpl implement interface LocationHandlerService
type LocationHandlerServiceImpl struct {
	inputService in.InLocationService
}

// NewLocationHandlerService DI by constructor
func NewLocationHandlerServiceImpl(inputService in.InLocationService) *LocationHandlerServiceImpl {
	return &LocationHandlerServiceImpl{inputService: inputService}
}

// HandleCreateLocation implement interface LocationHandlerService
func (i *LocationHandlerServiceImpl) HandleCreateLocation(ctx *gin.Context) {
	var ctxReq context.Context = ctx.Request.Context()
	request, ok := checkBindJsonError[dtos.Locationrequest](ctx)
	if !ok {
		return
	}
	bsLocation, err := i.inputService.CreateLocation(ctxReq, mappers.ToBusinessLocation(request))
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}

	ctx.JSON(http.StatusCreated, mappers.ToLocationResponse(bsLocation))
}

// HandleCreateLocation implement interface LocationHandlerService
func (i *LocationHandlerServiceImpl) HandleGetLocationByID(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	var ctxReq context.Context = ctx.Request.Context()
	bsLocation, err := i.inputService.GetLocationByID(ctxReq, id)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}

	ctx.JSON(http.StatusOK, mappers.ToLocationResponse(bsLocation))
}

// HandleGetAllLocation implement interface LocationHandlerService
func (i *LocationHandlerServiceImpl) HandleGetAllLocation(ctx *gin.Context) {
	bsLocations, err := i.inputService.GetAllLocation(ctx.Request.Context())
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	var locationsResponses []dtos.LocationResponse = make([]dtos.LocationResponse, len(bsLocations))
	for i, v := range bsLocations {
		locationsResponses[i] = mappers.ToLocationResponse(v)
	}

	ctx.JSON(http.StatusOK, locationsResponses)
}
