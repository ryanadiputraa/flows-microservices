package transaction

import (
	"context"

	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/domain"
)

type Usecase interface {
	AddTransaction(ctx context.Context, userID string, dto domain.TransactionDTO) (*domain.Transaction, error)
	GetTransactionSummary(ctx context.Context, UserID string) (*domain.TransactionSummary, error)
	ListTransactions(ctx context.Context, userID, start, end string, size, page int) ([]domain.Transaction, *domain.Meta, error)
}
