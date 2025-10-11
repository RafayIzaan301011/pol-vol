package application

import (
	"context"
	"pvms/pkg/errors"
	"pvms/volunteer/domain"
	"pvms/volunteer/domain/volunteer"
)

type GetAllVolunteerRequest struct {
	Volunteer *VolunteerFilter
}

type VolunteerFilter struct {
	Age          *int
	Phone        *string
	Role         *int
	City         *string
	Constituency *string
	IsActive     *bool

	Limit  *int
	Offset *int
}

func (r *GetAllVolunteerRequest) validate() error {
	ctx := context.TODO()
	if r.Volunteer == nil {
		return nil
	}

	if r.Volunteer.Age != nil && *r.Volunteer.Age < 18 {
		return errors.BadRequest(ctx, "age must be at least 18")
	}

	if r.Volunteer.Phone != nil && *r.Volunteer.Phone == "" {
		return errors.BadRequest(ctx, "phone is required")
	}

	if r.Volunteer.Role != nil && (*r.Volunteer.Role < int(domain.RoleTier1) || *r.Volunteer.Role > int(domain.RoleAdmin)) {
		return errors.BadRequest(ctx, "invalid role")
	}

	if r.Volunteer.City != nil && *r.Volunteer.City == "" {
		return errors.BadRequest(ctx, "city is required")
	}

	if r.Volunteer.Constituency != nil && *r.Volunteer.Constituency == "" {
		return errors.BadRequest(ctx, "constituency is required")
	}

	if r.Volunteer.Limit != nil && *r.Volunteer.Limit < 1 {
		return errors.BadRequest(ctx, "limit must be at least 1")
	}

	if r.Volunteer.Offset != nil && *r.Volunteer.Offset < 0 {
		return errors.BadRequest(ctx, "offset must be at least 0")
	}

	return nil
}

type GetAllVolunteer func(ctx context.Context, req GetAllVolunteerRequest) ([]domain.Volunteer, error)

func NewGetAllVolunteer(volunteerRepo volunteer.Repository) GetAllVolunteer {
	return func(ctx context.Context, req GetAllVolunteerRequest) ([]domain.Volunteer, error) {

		err := req.validate()
		if err != nil {
			return nil, errors.InternalServerErrorf(ctx, "failed to validate get all volunteer request: %v", err)
		}

		volunteers, err := volunteerRepo.GetAll(ctx, req.Volunteer.toDomain())
		if err != nil {
			return nil, err
		}

		return volunteers, nil
	}
}

func (r *VolunteerFilter) toDomain() *domain.VolunteerFilter {
	if r == nil {
		return nil
	}

	filter := &domain.VolunteerFilter{}
	if r.Age != nil {
		filter.Age = r.Age
	}
	if r.Phone != nil {
		filter.Phone = r.Phone
	}
	if r.Role != nil {
		filter.Role = r.Role
	}
	if r.City != nil {
		filter.City = r.City
	}
	if r.Constituency != nil {
		filter.Constituency = r.Constituency
	}
	if r.IsActive != nil {
		filter.IsActive = r.IsActive
	}
	if r.Limit != nil {
		filter.Limit = r.Limit
	}
	if r.Offset != nil {
		filter.Offset = r.Offset
	}

	return filter
}
