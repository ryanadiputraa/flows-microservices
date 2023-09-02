package user

import (
	"context"

	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
)

type Usecase interface {
	Register(ctx context.Context, user *domain.UserDTO) (*domain.User, error)
}