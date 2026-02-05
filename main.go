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
	http.HandleFunc("/api/clear-todos", handlers.ClearTodosHandler)
	http.HandleFunc("/api/upload", handlers.UploadFileHandler)
	http.HandleFunc("/api/files", handlers.ListFilesHandler)
	http.HandleFunc("/api/download", handlers.DownloadFileHandler)
	// End API routes

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
