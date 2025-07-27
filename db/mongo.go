package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Println("env load error", err)
	}

	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Println("connection error", err)
	}

	err = mongoClient.Ping(context.TODO(), readpref.Primary())
		if err != nil {
		log.Println("ping failed", err)
	}
	log.Println("mongo connected")
}