package user

import (
	"context"

	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
