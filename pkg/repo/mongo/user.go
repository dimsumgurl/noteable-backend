package mongo

import (
	"context"
	"errors"

	"github.com/dimsumgurl/noteable-backend/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUserCollection creates a user collection
func (c *MongoClient) CreateUserCollection() error {
	// If collection does not exist, then create
	// col, err := c.Database.ListCollectionNames(context.Background(), nil)
	// if err != nil {
	// 	return err
	// }
	// for _, v := range col {
	// 	if v == "User" {
	// 		// return errors.New("Collection already exists")
	// 	}
	// }
	err := c.Database.CreateCollection(context.Background(), "User")
	if err != nil {
		return err
	}
	return nil
}

// CreateUser creates a new user
func (c *MongoClient) CreateUser(user *models.UserLogin) error {
	col := c.getUserCollection()
	_, err := col.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

// RetrieveUserByID returns a user by ID
func (c *MongoClient) RetrieveUserByID(id int64) (*models.UserAuth, error) {
	col := c.getUserCollection()
	var result *models.UserAuth
	resp := col.FindOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	if resp == nil {
		return nil, errors.New("Not found")
	}
	err := resp.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// RetrieveUserByEmail returns a user by email address
func (c *MongoClient) RetrieveUserByEmail(email string) (*models.UserAuth, error) {
	col := c.getUserCollection()
	var result *models.UserAuth
	resp := col.FindOne(context.Background(), bson.D{{Key: "emailaddress", Value: email}})
	if resp == nil {
		return nil, errors.New("Not found")
	}
	err := resp.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *MongoClient) Update() error {
	return nil
}

func (c *MongoClient) Delete() error {
	return nil
}

func (c *MongoClient) getUserCollection() *mongo.Collection {
	col := c.Database.Collection("User")
	return col
}
