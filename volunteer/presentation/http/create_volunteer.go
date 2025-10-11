package http

import (
	"net/http"
	"pvms/utils"
	"pvms/volunteer/application"
	"pvms/volunteer/domain"
	"pvms/volunteer/presentation/adapter"
	"pvms/volunteer/presentation/models"

	"github.com/gin-gonic/gin"
)

type CreateVolunteer gin.HandlerFunc

func NewVolunteer(createVolunteer application.CreateVolunteer) CreateVolunteer {
	return func(c *gin.Context) {

		ctx := c.Request.Context()
		var body models.CreateVolunteerRequest
		err := c.ShouldBindJSON(&body)
		if err != nil {
			mErr := utils.GetPresentationError(utils.ErrJSON(c))
			c.JSON(mErr.Status, mErr)
			return
		}

		// application layer dto
		req := application.CreateVolunteerRequest{
			Name:         body.Name,
			Age:          body.Age,
			Phone:        body.Phone,
			CNIC:         body.CNIC,
			City:         body.City,
			Constituency: body.Constituency,
			Role:         domain.Role(body.Role),
			IsActive:     body.Active,
		}

		// call the application layer function to create a volunteer
		volunteer, err := createVolunteer(ctx, req)
		if err != nil {
			mErr := utils.GetPresentationError(err)
			c.JSON(mErr.Status, mErr)
			return
		}

		// on succes return the volunteer response
		c.JSON(http.StatusOK, adapter.GetVolunteerResponseFromDomain(volunteer))
	}
}
