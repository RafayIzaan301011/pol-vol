package application

import (
	"context"
	"pvms/pkg/errors"
	"pvms/volunteer/domain"
	"pvms/volunteer/domain/volunteer"

	"github.com/google/uuid"
)

type CreateVolunteerRequest struct {
	Name         string
	Age          int
	Phone        string
	CNIC         string
	City         string
	Constituency string
	Role         domain.Role
	IsActive     bool
}

func (req *CreateVolunteerRequest) validate() error {

	ctx := context.TODO()
	if req.Name == "" {
		return errors.BadRequest(ctx, "name is required")
	}

	if req.Age < 18 {
		return errors.BadRequest(ctx, "age must be at least 18")
	}

	if req.Phone == "" {
		return errors.BadRequest(ctx, "phone is required")
	}

	if len(req.CNIC) != 13 {
		return errors.BadRequest(ctx, "cnic must be 13 characters long")
	}

	if req.City == "" {
		return errors.BadRequest(ctx, "city is required")
	}

	if req.Constituency == "" {
		return errors.BadRequest(ctx, "constituency is required")
	}

	if req.Role < domain.RoleTier1 || req.Role > domain.RoleAdmin {
		return errors.BadRequest(ctx, "invalid role")
	}

	return nil
}

type CreateVolunteer func(ctx context.Context, req CreateVolunteerRequest) (*domain.Volunteer, error)

func NewCreateVolunteer(volunteerRepo volunteer.Repository) CreateVolunteer {
	return func(ctx context.Context, req CreateVolunteerRequest) (*domain.Volunteer, error) {

		// validate the request
		err := req.validate()
		if err != nil {
			return nil, err
		}

		volID := uuid.New().String()

		// map to domain
		volunteer := domain.Volunteer{
			ID:           volID,
			Name:         req.Name,
			Age:          req.Age,
			Phone:        req.Phone,
			CNIC:         req.CNIC,
			City:         req.City,
			Constituency: req.Constituency,
			Role:         req.Role,
			IsActive:     req.IsActive,
		}

		savedVolunteer, err := volunteerRepo.Create(ctx, &volunteer)
		if err != nil {
			return nil, errors.InternalServerError(ctx, "failed to create volunteer")
		}

		return savedVolunteer, nil
	}
}
