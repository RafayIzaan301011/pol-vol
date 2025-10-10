package main

import (
	"net/http"

	"pvms/api/server"

	"github.com/gin-gonic/gin"
)

func main() {

	server.RunHTTPServer(func(router *gin.Engine) http.Handler {
		api := router.Group("/api")

		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "service healthy"})
		})

		// return router since gin.Engine implements http.Handler
		return router
	})
}
