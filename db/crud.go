package db

import (
	"context"
	"log"
	"time"

	"github.com/mostafaElsouifi/todo-app-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

//insert todo
func MongoInsertOne(todo model.Todo) *mongo.InsertOneResult {
	inserted, err := collection.InsertOne(ctx, todo)
	if err != nil {
		log.Fatal(err)
	}
	return inserted
}

//Get All todos
func MongoGetAllTodos() []primitive.M {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	var todos []primitive.M
	for cursor.Next(ctx) {
		var todo primitive.M
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}
	return todos
}

// update todo
func MongoUpdateTodo(id string, todo string) *mongo.UpdateResult {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	updatedTime := time.Now()
	newUpdate := bson.M{"$set": bson.M{"item": todo, "lastUpdated": updatedTime}}
	updated, err := collection.UpdateOne(ctx, bson.M{"_id": _id}, newUpdate)
	if err != nil {
		log.Fatal(err)
	}
	return updated
}

// Get one todo
func MongoGetOneTodo(id string) bson.M {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	resultCursor := collection.FindOne(ctx, bson.M{"_id": _id})
	var todo bson.M
	resultCursor.Decode(&todo)
	return todo

}

//delete one todo
func MongoDeleteOneTodo(id string) *mongo.DeleteResult {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	deletedTodo, err := collection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		log.Fatal(err)
	}
	return deletedTodo
}

//DELETE all todos
func MongoDeleteAllTodos() *mongo.DeleteResult {
	deletedResult, err := collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	return deletedResult
}

//TODO  Update  TODO
