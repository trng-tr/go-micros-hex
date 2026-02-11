package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/dtos"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/mappers"
)

const (
	fail             string = "FAIL"
	sucess           string = "SUCCESS"
	errInvalidParams string = "error: invalid parameter"
	errNotDigit      string = "error: invalid id value, must be a digit"
)

// buildOrderResponse util method of struct OrderHandlerImpl to build orderResponse
func (o *OrderHandlerImpl) buildOrderResponse(ctx *gin.Context, order domain.Order, lines []domain.OrderLine) (dtos.OrderResponse, error) {
	var orderLinesResponses []dtos.OrderLineResponse = make([]dtos.OrderLineResponse, 0, len(lines))
	var getReqCtx = ctx.Request.Context()
	for _, line := range lines {
		location, err := o.inRemoteLocation.GetRemoteLocationByID(ctx.Request.Context(), line.LocationID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
			return dtos.OrderResponse{}, err
		}
		product, err := o.inRemoteProdSvc.GetRemoteProductByID(getReqCtx, line.ProductID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
			return dtos.OrderResponse{}, err
		}
		orderLinesResponses = append(orderLinesResponses, mappers.ToOrderLineResponse(line, product, location))
	}

	customer, err := o.inRemoteCustomerSvc.GetRemoteCustomerByID(ctx.Request.Context(), order.CustomerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return dtos.OrderResponse{}, err
	}

	var orderResponse = dtos.OrderResponse{
		ID:                  order.ID,
		CustomerResponse:    mappers.ToCustomerResponse(customer),
		Status:              string(order.Status),
		OrderLinesResponses: orderLinesResponses,
		CreatedAt:           order.CreatedAt.Format(time.RFC3339),
	}

	return orderResponse, nil
}

// getID util function to get id from request parameter
func getID(ctx *gin.Context) (int64, bool) {
	var idStr string = ctx.Param("id")
	if strings.TrimSpace(idStr) == "" {
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(fail, errInvalidParams))
		return 0, false
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(fail, errNotDigit))
		return 0, false
	}

	return id, true
}
