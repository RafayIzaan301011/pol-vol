package http

import (
	"pvms/utils"
	"pvms/volunteer/presentation/models"

	"github.com/gin-gonic/gin"
)

type CreateVolunteer gin.HandlerFunc

func NewVolunteer() CreateVolunteer {
	return func(c *gin.Context) {
		var req models.CreateVolunteerRequest
		err := c.ShouldBindJSON(&req)
		if err != nil {
			mErr := utils.GetPresentationError(utils.ErrJSON(c))
			c.JSON(mErr.Status, mErr)
			return
		}

		// application layer dto

		// call the application layer function to create a volunteer

		// on succes return the volunteer id

	}
}
