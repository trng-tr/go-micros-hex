package routes

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	HandleCreateOrder(ctx *gin.Context)
	HandleGetOrderByID(ctx *gin.Context)
	HandleGetAllOrder(ctx *gin.Context)
	HandleDeleteOrder(ctx *gin.Context)
	HandleIncreaseOrderLineQuantity(ctx *gin.Context)
	HandleDecreaseOrderLineQuantity(ctx *gin.Context)
}
