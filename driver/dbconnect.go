package driver

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	gotenv.Load()
}
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
func Disconnect_DB(client *mongo.Client) {

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

}
