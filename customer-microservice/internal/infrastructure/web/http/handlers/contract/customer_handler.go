package contract

import "github.com/gin-gonic/gin"

// CustomerHandlderService interface that use gin-gonic/gin
type CustomerHandlerService interface {
	CreateCustomerHandler(c *gin.Context)
	GetCustomerHandlder(c *gin.Context)
	GetAllCustomersHandler(c *gin.Context)
	UpdateCustomerHandler(c *gin.Context)
	DeleteCustomerHandler(c *gin.Context)
}
