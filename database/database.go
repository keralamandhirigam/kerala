package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoConnection() (*mongo.Client, error) {
	opts := options.Client()
	opts.ApplyURI("mongodb+srv://andreamose988:E6xYOfUn0erPPQYa@cluster0.4wntr2x.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	} else {
		fmt.Println("connection success")
	}

	err = client.Ping(context.TODO(), readpref.PrimaryPreferred())
	if err != nil {
		return nil, err
	}

	return client, nil
}
