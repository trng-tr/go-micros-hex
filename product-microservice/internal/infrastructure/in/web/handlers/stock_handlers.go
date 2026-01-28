package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trng-tr/product-microservice/internal/application/ports/in"
	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/dtos"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/mappers"
)

// StockHandlerServiceImpl implement handlers interface
type StockHandlerServiceImpl struct {
	stckInPort     in.InStockService    //DI input port interface
	prodInPort     in.InProductService  //DI input port interface
	locationInPort in.InLocationService //DI input port interface
}

// NewStockHandlerServiceImpl injection par constructeur
func NewStockHandlerServiceImpl(stckInPort in.InStockService, prodInPort in.InProductService,
	locationInPort in.InLocationService) *StockHandlerServiceImpl {
	return &StockHandlerServiceImpl{
		stckInPort:     stckInPort,
		prodInPort:     prodInPort,
		locationInPort: locationInPort,
	}
}

// HandlerCreateStock implement interface
func (h *StockHandlerServiceImpl) HandleCreateStock(ctx *gin.Context) {
	webRequest, ok := checkBindJsonError[dtos.StockRequest](ctx)
	if !ok {
		return
	}
	var ctxRequest context.Context = ctx.Request.Context()
	bsStock, err := h.stckInPort.CreateStock(ctxRequest, mappers.ToBusinessStock(webRequest))
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctxRequest, bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsLocation, err := h.locationInPort.GetLocationByID(ctxRequest, bsStock.LocationID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}

	ctx.JSON(http.StatusCreated, buildStokResponse(bsStock, bsProduct, bsLocation))
}

// HandlerGetStockByID implement interface
func (h *StockHandlerServiceImpl) HandleGetStockByID(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	var ctxRequest context.Context = ctx.Request.Context()
	bsStock, err := h.stckInPort.GetStockByID(ctxRequest, id)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctxRequest, bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsLocation, err := h.locationInPort.GetLocationByID(ctxRequest, bsStock.LocationID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	ctx.JSON(http.StatusOK, buildStokResponse(bsStock, bsProduct, bsLocation))

}

// HandlerGetAllStocks implement interface
func (h *StockHandlerServiceImpl) HandleGetAllStocks(ctx *gin.Context) {
	bsStocks, err := h.stckInPort.GetAllStocks(ctx.Request.Context())
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	var stocksResponses = make([]dtos.StockResponse, 0, len(bsStocks))
	var ctxRequest context.Context = ctx.Request.Context()
	for _, stock := range bsStocks {
		bsProduct, err := h.prodInPort.GetProductByID(ctxRequest, stock.ProductID)
		if ok := checkInternalServerError(err, ctx); !ok {
			return
		}
		bsLocation, err := h.locationInPort.GetLocationByID(ctxRequest, stock.LocationID)
		if ok := checkInternalServerError(err, ctx); !ok {
			return
		}
		stocksResponses = append(stocksResponses, buildStokResponse(stock, bsProduct, bsLocation))
	}

	ctx.JSON(http.StatusOK, stocksResponses)
}

// HandlerSetStockQuantity implement interface
func (h *StockHandlerServiceImpl) HandleSetStockQuantity(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	quantityRequest, ok := checkBindJsonError[dtos.StockQuantityRequest](ctx)
	if !ok {
		return
	}
	var ctxRequest context.Context = ctx.Request.Context()
	bsStock, err := h.stckInPort.SetStockQuantity(ctxRequest, id, quantityRequest.Quantity)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctxRequest, bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsLocation, err := h.locationInPort.GetLocationByID(ctxRequest, bsStock.LocationID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	ctx.JSON(http.StatusOK, buildStokResponse(bsStock, bsProduct, bsLocation))
}

// HandlerIncreaseStockQuantity implement interface
func (h *StockHandlerServiceImpl) HandleIncreaseStockQuantity(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	quantityRequest, ok := checkBindJsonError[dtos.StockQuantityRequest](ctx)
	var ctxRequest context.Context = ctx.Request.Context()
	bsStock, err := h.stckInPort.IncreaseStockQuantity(ctxRequest, id, quantityRequest.Quantity)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctxRequest, bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsLocation, err := h.locationInPort.GetLocationByID(ctxRequest, bsStock.LocationID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	ctx.JSON(http.StatusOK, buildStokResponse(bsStock, bsProduct, bsLocation))
}

// HandlerDecreaseStockQuantity implement interface
func (h *StockHandlerServiceImpl) HandleDecreaseStockQuantity(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	var ctxRequest context.Context = ctx.Request.Context()
	quantityRequest, ok := checkBindJsonError[dtos.StockQuantityRequest](ctx)
	bsStock, err := h.stckInPort.DecreaseStockQuantity(ctxRequest, id, quantityRequest.Quantity)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctxRequest, bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsLocation, err := h.locationInPort.GetLocationByID(ctxRequest, bsStock.LocationID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	ctx.JSON(http.StatusOK, buildStokResponse(bsStock, bsProduct, bsLocation))
}

// HandleGetStockByProductID implement interface
func (h *StockHandlerServiceImpl) HandleGetStockByProductID(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	var ctxRequest context.Context = ctx.Request.Context()
	bsStock, err := h.stckInPort.GetStockByProductID(ctxRequest, id)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctxRequest, bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}

	bsLocation, err := h.locationInPort.GetLocationByID(ctxRequest, bsStock.LocationID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	ctx.JSON(http.StatusOK, buildStokResponse(bsStock, bsProduct, bsLocation))
}

func buildStokResponse(stock domain.Stock, product domain.Product, location domain.Location) dtos.StockResponse {
	return mappers.ToStockResponse(stock, product, location)
}
