package email

import (
	"context"

	"github.com/ryanadiputraa/flows/flows-microservices/notification/internal/domain"
)

type Usecase interface {
	RegisterNotification(ctx context.Context, dto domain.EmailDTO) error
}
