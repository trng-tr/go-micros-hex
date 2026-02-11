package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trng-tr/order-microservice/internal/application/in"
	"github.com/trng-tr/order-microservice/internal/domain"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/dtos"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/mappers"
)

// OrderHandlerImpl implement OrderHandler
type OrderHandlerImpl struct {
	inOrderSvc          in.InOrderService
	inOrderLineSvc      in.InOrderLineService
	inRemoteCustomerSvc in.RemoteCustomerService
	inRemoteProdSvc     in.RemoteProductService
	inRemoteLocation    in.RemoteLocationService
}

// NewOrderHandlerImpl DI by constuctor
func NewOrderHandlerImpl(order in.InOrderService, orderL in.InOrderLineService, rc in.RemoteCustomerService,
	rp in.RemoteProductService, inRemoteLocation in.RemoteLocationService) *OrderHandlerImpl {
	return &OrderHandlerImpl{
		inOrderSvc:          order,
		inOrderLineSvc:      orderL,
		inRemoteCustomerSvc: rc,
		inRemoteProdSvc:     rp,
		inRemoteLocation:    inRemoteLocation,
	}
}

// HandleCreateOrder implement interface OrderHandler
func (o *OrderHandlerImpl) HandleCreateOrder(ctx *gin.Context) {
	var request dtos.OrderRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}

	var bsLines = make([]domain.OrderLine, 0, len(request.OrderLines))
	for _, line := range request.OrderLines {
		bsLine := mappers.ToBusinessOrderLine(line)
		bsLines = append(bsLines, bsLine)
	}

	var ctxReq = ctx.Request.Context()

	savedOrder, err := o.inOrderSvc.CreateOrderWithOrderLines(ctxReq, request.CustomerID, bsLines)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}

	bsCustomer, err := o.inRemoteCustomerSvc.GetRemoteCustomerByID(ctxReq, savedOrder.CustomerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}

	var orderLinesResponses []dtos.OrderLineResponse
	for _, line := range savedOrder.Lines {
		location, err := o.inRemoteLocation.GetRemoteLocationByID(ctxReq, line.LocationID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
			return
		}
		product, err := o.inRemoteProdSvc.GetRemoteProductByID(ctxReq, line.ProductID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
			return
		}

		orderLinesResponses = append(orderLinesResponses, mappers.ToOrderLineResponse(line, product, location))
	}

	var orderResponse dtos.OrderResponse = dtos.OrderResponse{
		ID:                  savedOrder.ID,
		CustomerResponse:    mappers.ToCustomerResponse(bsCustomer),
		Status:              string(savedOrder.Status),
		OrderLinesResponses: orderLinesResponses,
		CreatedAt:           savedOrder.CreatedAt.Format(time.RFC3339),
	}

	ctx.JSON(http.StatusOK, orderResponse)
}

// HandleGetOrderByID implement interface OrderHandler
func (o *OrderHandlerImpl) HandleGetOrderByID(ctx *gin.Context) {

	id, ok := getID(ctx)
	if !ok {
		return
	}
	var getReqCtx = ctx.Request.Context()
	order, err := o.inOrderSvc.GetOrderByID(getReqCtx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}

	lines, err := o.inOrderLineSvc.GetOrderLinesByOrderID(getReqCtx, order.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
	}

	order.Lines = lines
	orderResponse, err := o.buildOrderResponse(ctx, order, order.Lines)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, orderResponse)
}

// HandleGetAllOrder implement interface OrderHandler
func (o *OrderHandlerImpl) HandleGetAllOrders(ctx *gin.Context) {
	var ctxReq = ctx.Request.Context()
	orders, err := o.inOrderSvc.GetAllOrders(ctxReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}

	var orderResponses []dtos.OrderResponse = make([]dtos.OrderResponse, 0, len(orders))
	for _, order := range orders {
		orderlines, err := o.inOrderLineSvc.GetOrderLinesByOrderID(ctxReq, order.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
			return
		}
		orderResponse, err := o.buildOrderResponse(ctx, order, orderlines)
		if err != nil {
			return
		}
		orderResponses = append(orderResponses, orderResponse)
	}
	ctx.JSON(http.StatusOK, orderResponses)
}

// // HandleGetAllOrder implement interface OrderHandler
func (o *OrderHandlerImpl) HandleDeleteOrder(ctx *gin.Context) {
	id, ok := getID(ctx)
	if !ok {
		return
	}
	err := o.inOrderSvc.DeleteOrder(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, dtos.NewResponse(sucess, "successfully deleted"))
}

// HandleIncreaseOrderLineQuantity implement interface OrderHandler
func (o *OrderHandlerImpl) HandleIncreaseOrderLineQuantity(ctx *gin.Context) {
	var request dtos.AjustQuantityRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}
	id, ok := getID(ctx)
	if !ok {
		return
	}
	var ctxReq = ctx.Request.Context()
	orderLine, err := o.inOrderLineSvc.IncreaseOrderLineQuantity(ctxReq, id, request.Quantity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	location, err := o.inRemoteLocation.GetRemoteLocationByID(ctxReq, orderLine.LocationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	product, err := o.inRemoteProdSvc.GetRemoteProductByID(ctxReq, orderLine.ProductID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
	}

	ctx.JSON(http.StatusAccepted, mappers.ToOrderLineResponse(orderLine, product, location))
}

// HandleDecreaseOrderLineQuantity implement interface OrderHandler
func (o *OrderHandlerImpl) HandleDecreaseOrderLineQuantity(ctx *gin.Context) {
	var request dtos.AjustQuantityRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}
	id, ok := getID(ctx)
	if !ok {
		return
	}
	var ctxReq = ctx.Request.Context()
	line, err := o.inOrderLineSvc.DecreaseOrderLineQuantity(ctxReq, id, request.Quantity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	location, err := o.inRemoteLocation.GetRemoteLocationByID(ctxReq, line.LocationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	product, err := o.inRemoteProdSvc.GetRemoteProductByID(ctxReq, line.ProductID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
	}

	ctx.JSON(http.StatusAccepted, mappers.ToOrderLineResponse(line, product, location))
}
