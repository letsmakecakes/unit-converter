package handlers

import (
	"converters"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func ConvertLength(c *gin.Context) {
	var requestData converters.LengthConverter

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Infof("Obtained request: %+v", requestData)

	result, err := converters.GetResult(requestData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use logrus for logging the result
	log.Infof("Computed result: %v", result)

	responseData := ResponseData{
		ComputedResult: result,
	}

	c.JSON(http.StatusOK, responseData)
}
