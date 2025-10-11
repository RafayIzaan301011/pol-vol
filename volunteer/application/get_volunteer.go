package application

import (
	"context"
	"pvms/pkg/errors"
	"pvms/volunteer/domain"
	"pvms/volunteer/domain/volunteer"
)

type GetVolunteerRequest struct {
	ID string
}

// validate function
func (r *GetVolunteerRequest) validate() error {
	ctx := context.TODO()
	if r.ID == "" {
		return errors.BadRequest(ctx, "invalid volunteer id")
	}
	return nil
}

type GetVolunteer func(ctx context.Context, req GetVolunteerRequest) (*domain.Volunteer, error)

func NewGetVolunteer(volunteerRepo volunteer.Repository) GetVolunteer {
	return func(ctx context.Context, req GetVolunteerRequest) (*domain.Volunteer, error) {

		err := req.validate()
		if err != nil {
			return nil, errors.InternalServerErrorf(ctx, "failed to validate get volunteer request: %v", err)
		}

		volunteer, err := volunteerRepo.Get(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return volunteer, nil
	}
}
