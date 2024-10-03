package main

import (
	"handlers"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logrus logger
	log := logrus.New()
	log.Out = os.Stdout
	log.SetLevel(logrus.InfoLevel)

	// Create Gin router
	router := gin.Default()

	// Custom logger middleware
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logEntry := logrus.WithFields(logrus.Fields{
			"client_ip":   param.ClientIP,
			"status_code": param.StatusCode,
			"method":      param.Method,
			"path":        param.Path,
			"latency":     param.Latency,
			"time":        time.Now().Format(time.RFC3339),
		})

		// Log request details
		logEntry.Info("Incoming request")

		// Return empty string to avoid duplicate logging
		return ""
	}))

	// Recovery middleware
	router.Use(gin.Recovery())

	// Static files and templates
	router.Static("/static", "../../web/static")
	router.LoadHTMLGlob("../../web/templates/*.html")

	// Home route
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Conversion route
	router.POST("/length", handlers.ConvertLength)
	router.POST("/weight", handlers.ConvertWeight)

	// Port for the server
	port := "3000"
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
