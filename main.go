package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mostafaElsouifi/todo-app-go/db"
	"github.com/mostafaElsouifi/todo-app-go/router"
)

func main() {
	fmt.Println("Welcome to TODO-APP from GO")
	r := router.Router()
	db.Init()
	log.Fatal(http.ListenAndServe(":3000", r))

}
