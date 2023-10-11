package repository

import (
	"context"
	"errors"
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

func (r *repository) Save(ctx context.Context, transaction domain.Transaction) error {
	q := `
		INSERT INTO transactions (id, user_id, title, description, amount, date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.DB.ExecContext(ctx, q, transaction.ID, transaction.UserID, transaction.Title,
		transaction.Description, transaction.Amount, transaction.Date, transaction.CreatedAt, transaction.UpdatedAt)

	return err
}

func (r *repository) List(ctx context.Context, userID string, start, end time.Time, size, page int) ([]domain.Transaction, error) {
	var q string
	if size < 1 || page < 1 {
		return nil, errors.New("query error: 'size' or 'page' must be greater than 1")
	}

	q = `
		SELECT COUNT(id) OVER() as total_data, id, title, description, amount, date FROM transactions 
		WHERE user_id = $1 AND date BETWEEN $4 AND $5
		ORDER BY date DESC LIMIT $2 OFFSET $3
	`

	transactions := []domain.Transaction{}
	offset := (page - 1) * size
	err := r.DB.Select(&transactions, q, userID, size, offset, start, end)

	return transactions, err
}
