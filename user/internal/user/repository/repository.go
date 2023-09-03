package repository

import (
	"context"

	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db     *mongo.Client
	dbName string
}

func NewRepository(db *mongo.Client, dbName string) user.Repository {
	return &repository{db: db, dbName: dbName}
}

func (r *repository) Save(ctx context.Context, user *domain.User) error {
	c := r.db.Database(r.dbName).Collection("users")
	_, err := c.InsertOne(ctx, user)
	return err
}
