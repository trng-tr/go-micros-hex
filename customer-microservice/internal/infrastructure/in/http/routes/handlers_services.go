package routes

/*en Go, les interfaces sont du côte de celui qui les utilise(DI)
et non pas du coté de cui qui les implémente, ici c'est la route
qui les injecte
*/
import "github.com/gin-gonic/gin"

type AddressHandlerService interface {
	AddressHandleCreate(c *gin.Context)
	AddressHandleGetById(c *gin.Context)
	AddressHandleGetAll(c *gin.Context)
	AddressHandleDelete(c *gin.Context)
}

// CustomerHandlderService interface that use gin-gonic/gin
type CustomerHandlerService interface {
	CustomerHandleCreate(c *gin.Context)
	CustomerHandleGetById(c *gin.Context)
	CustomerHandleGetAll(c *gin.Context)
	CustomerHandlePatch(c *gin.Context)
	CustomerHandleDelete(c *gin.Context)
}
