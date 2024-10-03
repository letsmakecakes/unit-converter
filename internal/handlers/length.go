package handlers

import (
	"converters"
	"models"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func ConvertLength(c *gin.Context) {
	var requestData models.Request

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Infof("obtained request: %+v", requestData)

	result, err := converters.ConvertLength(requestData.FromUnit, requestData.ToUnit, requestData.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use logrus for logging the result
	log.Infof("computed result: %+v", result)

	responseData := models.Response{
		ComputedResult: result,
	}

	c.JSON(http.StatusOK, responseData)
}
