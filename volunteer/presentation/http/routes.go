package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Register volunteer routes
	// todo: load the config

	router.Group("/volunteers")
}
