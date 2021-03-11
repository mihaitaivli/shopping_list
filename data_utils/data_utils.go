package data_utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateClient returns a pointer to a client
func CreateClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.ClientOptions{
		Auth: &options.Credential{
			Username: "localUser",
			Password: "localPassword",
		},
	}
	clientOptions.ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, &clientOptions)
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return client
}
