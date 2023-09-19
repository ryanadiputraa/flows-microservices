package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/transaction/internal/transaction"
)

type repository struct {
	DB *sqlx.DB
}

func NewRepository(DB *sqlx.DB) transaction.Repository {
	return &repository{DB: DB}
}

func (r *repository) Save(ctx context.Context, transaction *domain.Transaction) error {
	return nil
}

func (r *repository) List(ctx context.Context, start, end time.Time, size, page int) ([]domain.Transaction, error) {
	return nil, nil
}
