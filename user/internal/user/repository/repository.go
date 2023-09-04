package repository

import (
	"context"

	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/domain"
	"github.com/ryanadiputraa/flows/flows-microservices/user/internal/user"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *repository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user *domain.User
	c := r.db.Database(r.dbName).Collection("users")

	if err := c.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}
