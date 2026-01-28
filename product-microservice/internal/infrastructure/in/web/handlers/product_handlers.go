package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/trng-tr/product-microservice/internal/application/ports/in"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/dtos"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/mappers"
)

// ProductHandlerServiceImpl implement handlers interface
type ProductHandlerServiceImpl struct {
	inputPort in.InProductService //DI input port interface
}

// NewProductHandlerServiceImpl injection par constructeur
func NewProductHandlerServiceImpl(inputPort in.InProductService) *ProductHandlerServiceImpl {
	return &ProductHandlerServiceImpl{inputPort: inputPort}
}

// HandleSaveProduct implement ProductHandlerService
func (h *ProductHandlerServiceImpl) HandleSaveProduct(ctx *gin.Context) {

	request, ok := checkBindJsonError[dtos.ProductRequest](ctx)
	if !ok {
		return
	}
	product, err := h.inputPort.SaveProduct(ctx.Request.Context(), mappers.ToBusinessProduct(request))
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}

	ctx.JSON(http.StatusCreated, mappers.ToProductResponse(product))
}

// HandleGetProductByID implement ProductHandlerService
func (h *ProductHandlerServiceImpl) HandleGetProductByID(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	product, err := h.inputPort.GetProductByID(ctx.Request.Context(), id)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	ctx.JSON(http.StatusOK, mappers.ToProductResponse(product))
}

// HandleGetAllProducts implement ProductHandlerService
func (h *ProductHandlerServiceImpl) HandleGetAllProducts(ctx *gin.Context) {
	products, err := h.inputPort.GetAllProducts(ctx.Request.Context())
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	var responses = make([]dtos.ProductResponse, 0, len(products))
	for _, prod := range products {
		responses = append(responses, mappers.ToProductResponse(prod))
	}

	ctx.JSON(http.StatusOK, responses)
}

// HandlePatchProduct implement ProductHandlerService
func (h *ProductHandlerServiceImpl) HandlePatchProduct(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}

	webPatchRequest, ok := checkBindJsonError[dtos.ProductPatchRequest](ctx)
	if !ok {
		return
	}
	businessPatch := mappers.ToDomainPatchRequest(webPatchRequest)
	product, err := h.inputPort.PatchProduct(ctx.Request.Context(), id, businessPatch)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}

	ctx.JSON(http.StatusOK, mappers.ToProductResponse(product))
}

// HandleDeleteProduct implement ProductHandlerService
func (h *ProductHandlerServiceImpl) HandleDeleteProduct(ctx *gin.Context) {
	id, ok := getId(ctx)
	if !ok {
		return
	}
	err := h.inputPort.DeleteProduct(ctx.Request.Context(), id)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	respose := dtos.NewResponse(dtos.Success, "Successfully deleted")
	ctx.JSON(http.StatusOK, respose)
}

// HandleGetProductBySku implement ProductHandlerService
func (h *ProductHandlerServiceImpl) HandleGetProductBySku(ctx *gin.Context) {
	var sku = ctx.Param("sku")
	if strings.TrimSpace(sku) == "" {
		ctx.JSON(http.StatusBadRequest, errInputEmptySku)
	}
	product, err := h.inputPort.GetProductBySku(ctx.Request.Context(), sku)
	if ok := checkInternalServerError(err, ctx); !ok {
		return
	}
	rep := mappers.ToProductResponse(product)
	ctx.JSON(http.StatusOK, rep)
}
