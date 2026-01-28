package routes

/*en Go, les interfaces sont du côte de celui qui les utilise(DI)
et non pas du coté de cui qui les implémente, ici c'est la route
qui les injecte
*/
import "github.com/gin-gonic/gin"

type ProductHandlerService interface {
	HandleSaveProduct(ctx *gin.Context)
	HandleGetProductByID(ctx *gin.Context)
	HandleGetAllProducts(ctx *gin.Context)
	HandlePatchProduct(ctx *gin.Context)
	HandleDeleteProduct(ctx *gin.Context)
	HandleGetProductBySku(ctx *gin.Context)
}

type LocationHandlerService interface {
	HandleCreateLocation(ctx *gin.Context)
	HandleGetLocationByID(ctx *gin.Context)
	HandleGetAllLocation(ctx *gin.Context)
}

type StockHanderService interface {
	HandleCreateStock(ctx *gin.Context)
	HandleGetStockByID(ctx *gin.Context)
	HandleGetAllStocks(ctx *gin.Context)
	HandleSetStockQuantity(ctx *gin.Context)
	HandleIncreaseStockQuantity(ctx *gin.Context)
	HandleDecreaseStockQuantity(ctx *gin.Context)
	HandleGetStockByProductID(ctx *gin.Context)
}
