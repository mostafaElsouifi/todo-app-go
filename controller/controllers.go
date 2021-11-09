package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mostafaElsouifi/todo-app-go/db"
	"github.com/mostafaElsouifi/todo-app-go/model"
)

func AddNewTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo model.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	// check if there is no  a todo
	if todo.IsEmpty() {
		json.NewEncoder(w).Encode("Please write a todo")
		return
	}
	todo.Time = time.Now()
	// insert todo to database
	db.MongoInsertOne(todo)
	json.NewEncoder(w).Encode("New todo Added")
}
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos := db.MongoGetAllTodos()
	json.NewEncoder(w).Encode(todos)
}

func GetOneTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	todo := db.MongoGetOneTodo(params["id"])
	json.NewEncoder(w).Encode(todo)
}

func EditTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var newTodo model.Todo
	json.NewDecoder(r.Body).Decode(&newTodo)
	updatedTodo := db.MongoUpdateTodo(params["id"], newTodo.Item)

	json.NewEncoder(w).Encode(updatedTodo)
}
func DeleteOneTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deletedTodo := db.MongoDeleteOneTodo(params["id"])
	if deletedTodo.DeletedCount == 0 {
		json.NewEncoder(w).Encode("no todo with id : " + params["id"])
		return
	}
	json.NewEncoder(w).Encode("successfully deleted")
}
func DeleteAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deletedResult := db.MongoDeleteAllTodos()
	if deletedResult.DeletedCount == 0 {
		json.NewEncoder(w).Encode("No data to delete")
		return
	}
	json.NewEncoder(w).Encode("successfully deleted " + strconv.FormatInt(deletedResult.DeletedCount, 10))
}
