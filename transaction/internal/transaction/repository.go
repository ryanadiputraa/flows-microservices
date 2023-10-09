package transaction

import (
	"context"
	"time"

	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, transaction domain.Transaction) error
	List(ctx context.Context, userID string, start, end time.Time, size, page int) ([]domain.Transaction, error)
}
