package volunteer

import (
	"context"
	"pvms/volunteer/domain"
)

type Repository interface {
	ReadRepository
	WriteRepository
}

type ReadRepository interface {
	Get(ctx context.Context, ID int) (*domain.Volunteer, error)
	GetAll(ctx context.Context, volunteer *domain.Volunteer) ([]domain.Volunteer, error)
}

type WriteRepository interface {
	Create(ctx context.Context, volunteer *domain.Volunteer) (*domain.Volunteer, error)
	Update(ctx context.Context, volunteer *domain.Volunteer) (*domain.Volunteer, error)
}
