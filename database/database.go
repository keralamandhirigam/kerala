package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoConnection() (*mongo.Client, error) {
	opts := options.Client()
	opts.ApplyURI("mongodb://localhost:27017")

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.PrimaryPreferred())
	if err != nil {
		return nil, err
	}

	return client, nil
}
