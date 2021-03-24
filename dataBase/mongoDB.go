package mongoDB

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoClient is the variable of mongo.client
var MongoClient *mongo.Client

//InitRun function initialize the database
func InitRun() {
	MongoClient = initDB()
}

func initDB() *mongo.Client {
	fmt.Println("Connect to MongoDB")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://root:root@cluster0.yw2ix.mongodb.net/test"))

	if err != nil {
		fmt.Println("connect error!")
		log.Fatal(err)
	}
	return client
}
