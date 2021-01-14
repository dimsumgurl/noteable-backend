package examples

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	FirstName string `json: "firstName"`
	LastName  string
	Gender    string
}

func createConnection() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017/?connect=direct"))
	if err != nil {
		fmt.Printf("cant init client")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("cant connect client")
	}

	yua := &Person{
		FirstName: "Yua",
	}
	context := context.Background()
	person := client.Database("mongo").Collection("people")
	data, err := bson.Marshal(yua)
	if err != nil {
		fmt.Print("could not marshal data")
	}
	_, err = person.InsertOne(context, data)
	defer cancel()
	if err != nil {
		fmt.Printf("error inserting")
		fmt.Print(err)
	}
	// id := res.InsertedID
	// var results []Person
	// err = col.Find(yua).All(&results)

	// if err != nil {
	// 	fmt.Printf("things happened, faiuled")
	// 	fmt.Print(err)
	// }
	// fmt.Print(results)
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
