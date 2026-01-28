package routes

import "github.com/gin-gonic/gin"

type RouteRegistration struct {
	handler OrderHandler
}

func NewRouteRegistration(handler OrderHandler) *RouteRegistration {
	return &RouteRegistration{handler: handler}
}

func (rr *RouteRegistration) RegisterRoutes() *gin.Engine {
	engine := gin.Default()

	api := engine.Group("/api/v1")
	api.POST("/orders", rr.handler.HandleCreateOrder)
	api.GET("/orders", rr.handler.HandleGetAllOrder)
	api.GET("/orders/:id", rr.handler.HandleGetOrderByID)
	api.DELETE("/orders/:id", rr.handler.HandleDeleteOrder)
	api.PUT("/orderlines/increase/:id", rr.handler.HandleIncreaseOrderLineQuantity)
	api.PUT("/orderlines/decrease/:id", rr.handler.HandleDecreaseOrderLineQuantity)
	return engine
}
