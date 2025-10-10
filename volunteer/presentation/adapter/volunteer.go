package adapter

import (
	"pvms/volunteer/domain"
	"pvms/volunteer/presentation/models"
)

func GetVolunteerResponseFromDomain(vd *domain.Volunteer) *models.VolunteerResponse {
	return &models.VolunteerResponse{
		ID:           vd.ID,
		Name:         vd.Name,
		Age:          vd.Age,
		Phone:        vd.Phone,
		CNIC:         vd.CNIC,
		City:         vd.City,
		Constituency: vd.Constituency,
		Role:         vd.Role.String(),
		Active:       vd.IsActive,
	}
}
