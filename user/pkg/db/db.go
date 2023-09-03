package db

import (
	"context"

	"github.com/ryanadiputraa/flows/flows-microservices/user/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewDB(config config.DB) (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.DSN))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	uniqueUserEmailIndexModel(client, config.DB_Name)

	return client, nil
}

func uniqueUserEmailIndexModel(client *mongo.Client, db string) error {
	c := client.Database(db).Collection("users")
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := c.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}
