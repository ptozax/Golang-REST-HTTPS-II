package driver

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect_db() *mongo.Client {
	clientOptions := options.Client().
		ApplyURI(os.Getenv("User_DB"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GET_Mongo_Collection(collection_name string) (*mongo.Client, *mongo.Collection, *context.Context) {
	client := Connect_db()
	collection := client.Database(os.Getenv("DB_name")).Collection(collection_name)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return client, collection, &ctx
}
