package database

import (
	"context"
	"log"
	"processlist/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

var Db *mongo.Database

func CreateClient() {

	Client, err := mongo.NewClient(options.Client().ApplyURI(configs.MONGO_HOST))

	if err != nil {
		log.Println("[ERROR] cound not create client for database", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	err = Client.Connect(ctx)

	defer cancel()

	if err != nil {
		log.Println("[ERROR] cound not create client for database", err)
	}

	Db = Client.Database("ProcessList")

	return
}
