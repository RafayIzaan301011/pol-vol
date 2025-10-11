package http

import (
	"net/http"
	"pvms/utils"
	"pvms/volunteer/application"
	"pvms/volunteer/presentation/adapter"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetVolunteer gin.HandlerFunc

func NewGetVolunteer(getVolunteer application.GetVolunteer) GetVolunteer {
	return func(c *gin.Context) {

		ctx := c.Request.Context()
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			mErr := utils.GetPresentationError(utils.ErrUserID(ctx))
			c.JSON(mErr.Status, mErr)
			return
		}

		req := application.GetVolunteerRequest{
			ID: id.String(),
		}

		volunteer, err := getVolunteer(ctx, req)
		if err != nil {
			mErr := utils.GetPresentationError(err)
			c.JSON(mErr.Status, mErr)
			return
		}

		c.JSON(http.StatusOK, adapter.GetVolunteerResponseFromDomain(volunteer))
	}
}
