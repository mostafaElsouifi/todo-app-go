package db

import (
	"context"
	"fmt"
	"log"

	"github.com/mostafaElsouifi/todo-app-go/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var connectionUrl = "mongodb+srv://mostafa:" + env.GetEnv("password") + "@cluster0.hxecr.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
var dbName = "Todo-app"
var collectionName = "todos"

func Init() {

	clientOption := options.Client().ApplyURI(connectionUrl)
	client, er := mongo.Connect(context.TODO(), clientOption)
	if er != nil {
		log.Fatal(er)
	}
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Database connected")
}
