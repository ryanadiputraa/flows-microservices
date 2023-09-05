package user

import (
	"context"

	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
)

type Usecase interface {
	Register(ctx context.Context, user *domain.UserDTO) (*domain.User, error)
	Login(ctx context.Context, dto *domain.LoginDTO) (*domain.User, error)
	GetUserInfo(ctx context.Context, userID string) (*domain.User, error)
}
