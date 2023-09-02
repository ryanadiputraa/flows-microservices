package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) user.Repository {
	return &repository{db: db}
}

func (r *repository) Save(ctx context.Context, user *domain.User) error {
	q := `INSERT INTO users (id, email, password, first_name, last_name, currency, picture, created_at)
	VALUES (:id, :email, :password, :first_name, :last_name, :currency, :picture, :created_at)`

	_, err := r.db.NamedExecContext(ctx, q, user)

	return err
}
