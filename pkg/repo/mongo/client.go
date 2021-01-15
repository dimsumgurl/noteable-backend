package mongo

import (
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	ClientObject *mongo.Client
}

type ClientOptions struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewClient(clientOptions *ClientOptions) (*MongoClient, error) {
	port := strconv.Itoa(clientOptions.Port)
	hostURL := "mongodb://" + clientOptions.Username + ":" + clientOptions.Password + "@" + clientOptions.Host + ":" + port + "/?connect=direct"
	client, err := mongo.NewClient(options.Client().ApplyURI(hostURL))

	if err != nil {
		return nil, err
	}

	return &MongoClient{ClientObject: client}, nil
}

func (c *MongoClient) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := c.ClientObject.Connect(ctx)

	if err != nil {
		return err
	}

	defer func() {
		if err = c.ClientObject.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return err
}