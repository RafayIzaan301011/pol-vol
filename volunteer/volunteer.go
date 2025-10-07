package volunteerss

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// this file will intiate and register the volunteer routes
// this file will is used to expose the application functions to the other packages
// InitiateAndRegisterVolunteerRoutes calls the function to register volunteer routes in the routes.go file

func InitiateAndRegisterVolunteerRoutes(router *gin.Engine, db *gorm.DB) {

	//todo: load the config

	// group the volunteer routes

}
