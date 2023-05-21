package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBAPIWrapper struct {
	client *mongo.Client
}

func CreateMongoAPIWrapper() *MongoDBAPIWrapper {
	fmt.Println(os.Getenv("MONGODB_USER_PASSWORD"))
	connectionString := fmt.Sprintf(
		"mongodb+srv://pelee:%s@via.fn83vy3.mongodb.net/?retryWrites=true&w=majority",
		os.Getenv("MONGODB_USER_PASSWORD"),
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}

	wrapper := &MongoDBAPIWrapper{
		client: client,
	}

	return wrapper
}

func (wrapper *MongoDBAPIWrapper) CloseMongoDBAPIWrapper() {
	if wrapper.client == nil {
		return
	}

	err := wrapper.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}

func (w *MongoDBAPIWrapper) Ping() {
	if err := w.client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Ping")
}
