package routes

/*en Go, les interfaces sont du côte de celui qui les utilise(DI)
et non pas du coté de cui qui les implémente, ici c'est la route
qui les injecte
*/
import (
	"github.com/gin-gonic/gin"
)

// RoutesRegister struct to use interface contract.AddressHandlerService
type RoutesRegistration struct {
	addressHandler  AddressHandlerService
	customerHandler CustomerHandlerService
}

// NewRoutesRegistration DI par constructeur
func NewRoutesRegistration(addressHandler AddressHandlerService,
	customerHandler CustomerHandlerService) *RoutesRegistration {
	return &RoutesRegistration{
		addressHandler:  addressHandler,
		customerHandler: customerHandler,
	}
}

// RegisterRoutes method of struct NewRoutesRegistration
func (rr *RoutesRegistration) RegisterRoutes() *gin.Engine {
	//var engine = gin.Default() //gin.Default ajoute automatiqment les middlewre gin.Logger() et gin.Recovery()
	var engine *gin.Engine = gin.New()
	/*gin.Logger() middleware pour logger les request
	gin.Recovery middleware pour  eviter que le server tombe en cas de panic*/
	engine.Use(gin.Logger(), gin.Recovery())
	var api = engine.Group("/api/v1")

	api.POST("/addresses", rr.addressHandler.AddressHandleCreate)
	api.GET("/addresses/:id", rr.addressHandler.AddressHandleGetById)
	api.GET("/addresses", rr.addressHandler.AddressHandleGetAll)
	api.DELETE("/addresses/:id", rr.addressHandler.AddressHandleDelete)
	api.POST("/customers", rr.customerHandler.CustomerHandleCreate)
	api.GET("/customers/:id", rr.customerHandler.CustomerHandleGetById)
	api.GET("/customers", rr.customerHandler.CustomerHandleGetAll)
	api.PATCH("/customers/:id", rr.customerHandler.CustomerHandlePatch)
	api.DELETE("/customers/:id", rr.customerHandler.CustomerHandleDelete)

	return engine
}
