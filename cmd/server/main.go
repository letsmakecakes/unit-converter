package main

import (
	"handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "../../web/static")
	router.LoadHTMLGlob("../../web/templates/*.html")

	// Home route (index)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/length", handlers.ConvertLength)

	port := "3000"

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
