package repository

import (
	"context"
	"pvms/volunteer/domain"
)

type Repository interface {
}

type ReadRepository interface {
	Get(ctx context.Context, ID int) *domain.Volunteer
	GetAll(ctx context.Context, volunteer *domain.Volunteer) []domain.Volunteer
}

type WriteRepository interface {
	Create(ctx context.Context, volunteer *domain.Volunteer) int
	Update(ctx context.Context, volunteer *domain.Volunteer) *domain.Volunteer
}
