package contract

import "github.com/gin-gonic/gin"

type AddressHandlerService interface {
	CreateAddressHandler(c *gin.Context)
	GetAddressHandler(c *gin.Context)
}
