package application

import (
	"context"

	"pvms/pkg/errors"
	"pvms/volunteer/domain"
	"pvms/volunteer/domain/security"
	"pvms/volunteer/domain/volunteer"
)

type UpdateVolunteerRequest struct {
	UserID int
	ID     string
	Name   *string
	Age    *int

	Phone *string
	CNIC  *string

	Role         *int
	City         *string
	Constituency *string
	IsActive     *bool
}

func (r *UpdateVolunteerRequest) validate() error {

	ctx := context.TODO()

	if r.UserID == 0 {
		return errors.BadRequest(ctx, "user id is required")
	}

	if r.ID == "" {
		return errors.BadRequest(ctx, "id is required")
	}

	if r.Age != nil && *r.Age < 18 {
		return errors.BadRequest(ctx, "age must be at least 18")
	}

	if r.Phone != nil && *r.Phone == "" {
		return errors.BadRequest(ctx, "phone is required")
	}

	if r.CNIC != nil && len(*r.CNIC) != 13 {
		return errors.BadRequest(ctx, "cnic must be 13 characters long")
	}

	if r.Role != nil && (*r.Role < int(domain.RoleTier1) || *r.Role > int(domain.RoleAdmin)) {
		return errors.BadRequest(ctx, "invalid role")
	}

	return nil
}

func (r *UpdateVolunteerRequest) toDomain() *domain.Volunteer {
	volunteer := &domain.Volunteer{
		ID: r.ID,
	}

	if r.Name != nil {
		volunteer.Name = *r.Name
	}
	if r.Age != nil {
		volunteer.Age = *r.Age
	}
	if r.Phone != nil {
		volunteer.Phone = *r.Phone
	}
	if r.CNIC != nil {
		volunteer.CNIC = *r.CNIC
	}
	if r.Role != nil {
		volunteer.Role = domain.Role(*r.Role)
	}
	if r.City != nil {
		volunteer.City = *r.City
	}
	if r.Constituency != nil {
		volunteer.Constituency = *r.Constituency
	}
	if r.IsActive != nil {
		volunteer.IsActive = *r.IsActive
	}

	return volunteer
}

type UpdateVolunteer func(ctx context.Context, req UpdateVolunteerRequest) (*domain.Volunteer, error)

func NewUpdateVolunteer(volunteerRepo volunteer.Repository, volunteerSecurity security.Encryption) UpdateVolunteer {
	return func(ctx context.Context, req UpdateVolunteerRequest) (*domain.Volunteer, error) {
		if err := req.validate(); err != nil {
			return nil, err
		}

		volunteerDomain := req.toDomain()
		updatedVolunteer, err := volunteerRepo.Update(ctx, volunteerDomain)
		if err != nil {
			return nil, errors.InternalServerErrorf(ctx, "failed to update volunteer: %v", err)
		}

		return updatedVolunteer, nil
	}
}
