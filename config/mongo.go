package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("Mongo client creation failed: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Mongo connection failed: ", err)
	}

	DB = client.Database(os.Getenv("MONGO_DB_NAME"))
	log.Println("MongoDB connected")
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
