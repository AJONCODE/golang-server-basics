package server

import (
	"log"
	"net/http"

	"github.com/AJONCODE/golangServerMisc/01_server/route"
)

// Server with route
func Server() {
	http.HandleFunc("/home", route.HomeRoute())
	http.HandleFunc("/todo", route.GetTodos())
	http.HandleFunc("/add-todo", route.AddTodo())

	log.Println("Server is running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
