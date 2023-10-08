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
	q := `INSERT INTO transactions (id, user_id, title, description, amount, date, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.DB.ExecContext(ctx, q, transaction.ID, transaction.UserID, transaction.Title,
		transaction.Description, transaction.Amount, transaction.Date, transaction.CreatedAt, transaction.UpdatedAt)

	return err
}

func (r *repository) List(ctx context.Context, start, end time.Time, size, page int) ([]*domain.Transaction, error) {
	return nil, nil
}
