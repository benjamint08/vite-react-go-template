package main

import (
	"github.com/benjamint08/vite-react-go-template/handlers"
	"github.com/benjamint08/vite-react-go-template/helpers"
	"net/http"
)

func main() {
	helpers.CheckFlags()

	// Define API routes here
	http.HandleFunc("/api/hello", handlers.HelloHandler)
	http.HandleFunc("/api/todos", handlers.GetTodoHandler)
	http.HandleFunc("/api/delete-todo", handlers.DeleteTodoHandler)
	http.HandleFunc("/api/add-todo", handlers.AddTodoHandler)
	// End API routes

	http.ListenAndServe(":8080", nil)
}
