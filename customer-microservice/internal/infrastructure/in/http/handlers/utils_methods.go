package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/trng-tr/customer-microservice/internal/infrastructure/in/http/dtos"
)

const (
	fail    string = "Failed"
	success string = "Success"
)

var errMethodNotAllowed error = errors.New("method not allowed")

// checkPostRequest private utility function
func checkPostRequest(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, dtos.NewResponse(fail, errMethodNotAllowed.Error()))
		return
	}
}

// checkGetMethod private utility function
func checkGetMethod(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.JSON(http.StatusMethodNotAllowed, dtos.NewResponse(fail, errMethodNotAllowed.Error()))
		return
	}
}

func checkDeleteMethod(c *gin.Context) {
	if c.Request.Method != http.MethodDelete {
		c.JSON(http.StatusMethodNotAllowed, dtos.NewResponse(fail, errMethodNotAllowed.Error()))
		return
	}
}

// checkId private utility function
func getId(c *gin.Context) (int64, error) {
	var strId string = c.Param("id")
	if strings.TrimSpace(strId) == "" {
		return -1, errors.New("error: you have to input id")
	}

	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("error: you have to enter a digit for id")
	}

	return id, nil
}
