package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trng-tr/product-microservice/internal/application/ports/in"
	"github.com/trng-tr/product-microservice/internal/domain"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/dtos"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/mappers"
)

// StockHandlerServiceImpl implement handlers interface
type StockHandlerServiceImpl struct {
	stckInPort in.InStockService   //DI input port interface
	prodInPort in.InProductService //DI input port interface
}

// NewStockHandlerServiceImpl injection par constructeur
func NewStockHandlerServiceImpl(stckInPort in.InStockService, prodInPort in.InProductService) *StockHandlerServiceImpl {
	return &StockHandlerServiceImpl{stckInPort: stckInPort, prodInPort: prodInPort}
}

// HandlerCreateStock implement interface
func (h *StockHandlerServiceImpl) HandleCreateStock(ctx *gin.Context) {
	webRequest, ok := checkBindJsonError[dtos.StockRequest](ctx)
	if !ok {
		return
	}
	bsStock, err := h.stckInPort.CreateStock(ctx.Request.Context(), mappers.ToBusinessStock(webRequest))
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctx.Request.Context(), bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}

	ctx.JSON(http.StatusCreated, buildStokResponse(bsStock, bsProduct))
}

// HandlerGetStockByID implement interface
func (h *StockHandlerServiceImpl) HandleGetStockByID(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	bsStock, err := h.stckInPort.GetStockByID(ctx.Request.Context(), id)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctx.Request.Context(), bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}

	ctx.JSON(http.StatusOK, buildStokResponse(bsStock, bsProduct))

}

// HandlerGetAllStocks implement interface
func (h *StockHandlerServiceImpl) HandleGetAllStocks(ctx *gin.Context) {
	bsStocks, err := h.stckInPort.GetAllStocks(ctx.Request.Context())
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	var stocksResponses = make([]dtos.StockResponse, 0, len(bsStocks))
	for _, stock := range bsStocks {
		bsProduct, err := h.prodInPort.GetProductByID(ctx.Request.Context(), stock.ProductID)
		if ok := checkInternalServerError(err, ctx); !ok {
			return
		}
		stocksResponses = append(stocksResponses, buildStokResponse(stock, bsProduct))
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
	bsStock, err := h.stckInPort.SetStockQuantity(ctx.Request.Context(), id, quantityRequest.Quantity)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctx.Request.Context(), bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	ctx.JSON(http.StatusOK, buildStokResponse(bsStock, bsProduct))
}

// HandlerIncreaseStockQuantity implement interface
func (h *StockHandlerServiceImpl) HandleIncreaseStockQuantity(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	quantityRequest, ok := checkBindJsonError[dtos.StockQuantityRequest](ctx)
	bsStock, err := h.stckInPort.IncreaseStockQuantity(ctx.Request.Context(), id, quantityRequest.Quantity)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctx.Request.Context(), bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	ctx.JSON(http.StatusOK, buildStokResponse(bsStock, bsProduct))
}

// HandlerDecreaseStockQuantity implement interface
func (h *StockHandlerServiceImpl) HandleDecreaseStockQuantity(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	quantityRequest, ok := checkBindJsonError[dtos.StockQuantityRequest](ctx)
	bsStock, err := h.stckInPort.DecreaseStockQuantity(ctx.Request.Context(), id, quantityRequest.Quantity)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctx.Request.Context(), bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	ctx.JSON(http.StatusOK, buildStokResponse(bsStock, bsProduct))
}

// HandleGetStockByProductID implement interface
func (h *StockHandlerServiceImpl) HandleGetStockByProductID(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	bsStock, err := h.stckInPort.GetStockByProductID(ctx.Request.Context(), id)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	bsProduct, err := h.prodInPort.GetProductByID(ctx.Request.Context(), bsStock.ProductID)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}

	ctx.JSON(http.StatusOK, buildStokResponse(bsStock, bsProduct))
}

func buildStokResponse(stock domain.Stock, product domain.Product) dtos.StockResponse {
	return mappers.ToStockResponse(stock, product)
}
