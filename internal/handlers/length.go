package handlers

import (
	"converters"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConvertLength(c *gin.Context) {
	var requestData converters.LengthConverter

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := converters.GetResult(requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	responseData := ResponseData{
		ComputedResult: result,
	}

	c.JSON(http.StatusOK, responseData)
}
