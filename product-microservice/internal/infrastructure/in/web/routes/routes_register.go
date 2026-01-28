package routes

/*en Go, les interfaces sont du côte de celui qui les utilise(DI)
et non pas du coté de cui qui les implémente, ici c'est la route
qui les injecte
*/
import "github.com/gin-gonic/gin"

//Routes injecte by DI ProductHandlerService and StockHanderService
type Routes struct {
	phs ProductHandlerService  //DI
	shs StockHanderService     //DI
	lhs LocationHandlerService //DI
}

//NewRoutes DI by constructor
func NewRoutes(phs ProductHandlerService, shs StockHanderService, lhs LocationHandlerService) *Routes {
	return &Routes{phs: phs, shs: shs, lhs: lhs}
}

//RegisterRoutes func method
func (r *Routes) RegisterApiRoutes() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())
	api := engine.Group("/api/v1")
	//for product service
	api.POST("/products", r.phs.HandleSaveProduct)
	api.GET("/products", r.phs.HandleGetAllProducts)
	api.GET("/products/:id", r.phs.HandleGetProductByID)
	api.GET("/products/sku/:sku", r.phs.HandleGetProductBySku)
	api.PATCH("/products/:id", r.phs.HandlePatchProduct)
	api.DELETE("/products/:id", r.phs.HandleDeleteProduct)

	//for location service
	api.POST("/locations", r.lhs.HandleCreateLocation)
	api.GET("/locations", r.lhs.HandleGetAllLocation)
	api.GET("/locations/:id", r.lhs.HandleGetLocationByID)

	//for stock service
	api.POST("/stocks", r.shs.HandleCreateStock)
	api.GET("/stocks", r.shs.HandleGetAllStocks)
	api.GET("/stocks/:id", r.shs.HandleGetStockByID)
	api.PUT("/stocks/set-qte/:id", r.shs.HandleSetStockQuantity)
	api.PUT("/stocks/increase-qte/:id", r.shs.HandleIncreaseStockQuantity)
	api.PUT("/stocks/decrease-qte/:id", r.shs.HandleDecreaseStockQuantity)
	api.GET("/products/:id/stock", r.shs.HandleGetStockByProductID)
	return engine
}
