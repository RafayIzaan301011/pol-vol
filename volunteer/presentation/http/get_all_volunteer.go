package http

import (
	"net/http"
	"pvms/utils"
	"pvms/volunteer/application"
	"pvms/volunteer/presentation/adapter"

	"github.com/gin-gonic/gin"
)

type GetAllVolunteer gin.HandlerFunc

func NewGetAllVolunteer(getAllVolunteer application.GetAllVolunteer) GetAllVolunteer {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// age, err := utils.ParseQueryParamToInt(c.Query("age"))
		// if err != nil {
		// 	mErr := utils.GetPresentationError(err)
		// 	c.JSON(mErr.Status, mErr)
		// 	return
		// }

		req := application.GetAllVolunteerRequest{
			Volunteer: application.VolunteerFilter{},
		}

		volunteer, err := getAllVolunteer(ctx, req)
		if err != nil {
			mErr := utils.GetPresentationError(err)
			c.JSON(mErr.Status, mErr)
			return
		}

		c.JSON(http.StatusOK, adapter.GetAllVolunteerResponseFromDomain(volunteer))
	}
}

// age
