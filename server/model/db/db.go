package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseClient struct {
	Client *mongo.Client
	Ctx    *context.Context
}

func (app *DatabaseClient) GetCollection(collectionName string) (*mongo.Collection, error) {
	collections, err := app.Client.Database("bank").ListCollectionNames(*app.Ctx, bson.D{})

	if err != nil {
		return nil, err
	}
	for iter := range collections {
		if collections[iter] == collectionName {

			return app.Client.Database("bank").Collection(collectionName), nil
		}
	}
	return nil, errors.New("not found")
}
