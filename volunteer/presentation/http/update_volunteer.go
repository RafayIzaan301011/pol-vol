package http

import (
	"net/http"
	"pvms/utils"
	"pvms/volunteer/application"
	"pvms/volunteer/presentation/adapter"
	"pvms/volunteer/presentation/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateVonteer gin.HandlerFunc

func NewUpdateVolunteer(updateVolunteer application.UpdateVolunteer) UpdateVonteer {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		Id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			mErr := utils.GetPresentationError(ErrUserID(ctx))
			c.JSON(mErr.Status, mErr)
		}

		var body models.UpdateVolunteerRequest
		err = c.ShouldBindJSON(&body)
		if err != nil {
			if err != nil {
				mErr := utils.GetPresentationError(ErrJSON(ctx))
				c.JSON(mErr.Status, mErr)
			}
		}

		req := application.UpdateVolunteerRequest{
			ID:           Id.String(),
			Name:         body.Name,
			Age:          body.Age,
			Phone:        body.Phone,
			CNIC:         body.CNIC,
			City:         body.City,
			Constituency: body.Constituency,
			Role:         body.Role,
			IsActive:     body.Active,
		}

		updatedVolunteer, err := updateVolunteer(ctx, req)
		if err != nil {
			mErr := utils.GetPresentationError(err)
			c.JSON(mErr.Status, mErr)
			return
		}

		c.JSON(http.StatusOK, adapter.GetVolunteerResponseFromDomain(updatedVolunteer))

	}
}
