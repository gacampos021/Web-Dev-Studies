package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://gacampos021_db_user:0O3xyW5iowAFP2eT@cluster0.nm3kjap.mongodb.net/?appName=Cluster0"

var MongoClient *mongo.Client

func init() {
	if err := connect_mongodb(); err != nil {
		log.Fatal("error connecting mongodb")
	}
}

func connect_mongodb() error {
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
