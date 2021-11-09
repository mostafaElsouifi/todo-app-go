package router

import (
	"github.com/gorilla/mux"
	"github.com/mostafaElsouifi/todo-app-go/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/todo", controller.AddNewTodo).Methods("POST")
	router.HandleFunc("/api/todo/{id}", controller.GetOneTodo).Methods("GET")
	router.HandleFunc("/api/todo/{id}", controller.EditTodo).Methods("PUT")
	router.HandleFunc("/api/todo/{id}", controller.DeleteOneTodo).Methods("DELETE")
	router.HandleFunc("/api/todos", controller.GetAllTodos).Methods("GET")
	router.HandleFunc("/api/todos", controller.DeleteAllTodos).Methods("DELETE")
	return router
}
