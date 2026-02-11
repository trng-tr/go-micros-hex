package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/trng-tr/product-microservice/internal/infrastructure/in/web/dtos"
)

var (
	errInputEmptyId   error = errors.New("error: input id is empty")
	errInputEmptySku  error = errors.New("error: input sku is empty")
	errInputNotDigit  error = errors.New("error: input id is not a digit")
	errInvalidInputId error = errors.New("error: input id invalid")
)

// getId util func
func getId(ctx *gin.Context) (int64, bool) {
	var rawID = ctx.Param("id")
	rawID = strings.TrimSpace(rawID)
	if rawID == "" {
		log.Println(errInputEmptyId)
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(dtos.Fail, errInputEmptyId.Error()))
		return 0, false
	}
	id, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(dtos.Fail, errInputNotDigit.Error()))
		return 0, false
	}
	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(dtos.Fail, errInvalidInputId.Error()))
		return 0, false
	}
	return id, true
}

// checkInternalServerError util func
func checkInternalServerError(err error, ctx *gin.Context) bool {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(dtos.Fail, err.Error()))
		return false
	}

	return true
}

// checkBadRequestError util func
func checkBindJsonError[Request any](ctx *gin.Context) (Request, bool) {
	var request Request
	//get json payload and write it in request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		log.Println(err.Error())
		var Zero Request
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(dtos.Fail, err.Error()))
		return Zero, false
	}

	return request, true
}

func getStockByLocationAndProductIDs(ctx *gin.Context) (int64, int64, bool) {
	locationId, err := strconv.ParseInt(ctx.Param("locationId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(dtos.Fail, err.Error()))
		return 0, 0, false
	}
	productId, err := strconv.ParseInt(ctx.Param("productId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(dtos.Fail, err.Error()))
		return 0, 0, false
	}

	return locationId, productId, true
}
