package mongo

import (
	"context"
	"log"

	"study-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func init() {
	if err := connect_mongodb(); err != nil {
		log.Fatal("error connecting mongodb")
	}
}

func connect_mongodb() error {
	config.LoadEnv()
	uri := config.GetEnv("URI_MONGO")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions((serverAPI))

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	MongoClient = client
	return err
}
