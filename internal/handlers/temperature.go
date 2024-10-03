package handlers

import (
	"converters"
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ConvertTemperature(c *gin.Context) {
	var requestData models.Request

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Infof("obtained request: %+v", requestData)

	result, err := converters.ConvertTemperature(requestData.FromUnit, requestData.ToUnit, requestData.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Infof("computed result: %+v", result)

	responseData := models.Response{
		ComputedResult: result,
	}

	c.JSON(http.StatusOK, responseData)
}
