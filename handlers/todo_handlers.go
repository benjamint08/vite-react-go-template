package handlers

import (
	"encoding/json"
	"git.ben.cymru/ben/mongohelper"
	"github.com/benjamint08/vite-react-go-template/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	todos := mongohelper.FindManyDocuments("vite-react-go", "todos", bson.D{{}})
	json.NewEncoder(w).Encode(todos)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body := r.Body
	defer body.Close()
	bodyJson := helpers.GetJsonFromBody(r)
	bodyMap, ok := bodyJson.(map[string]interface{})
	if !ok {
		http.Error(w, "Failed to decode body", http.StatusInternalServerError)
		return
	}
	if bodyMap["error"] != nil {
		http.Error(w, "Failed to decode body", http.StatusInternalServerError)
		return
	}

	todo := bodyMap["todo"].(string)
	exists := false
	todoExists := mongohelper.FindOneDocument("vite-react-go", "todos", bson.D{{"todo", todo}})
	if todoExists != nil {
		exists = true
	}
	if !exists {
		http.Error(w, "Todo does not exist", http.StatusBadRequest)
		return
	}
	deleteTodo := mongohelper.DeleteOneDocument("vite-react-go", "todos", bson.D{{"todo", todo}})
	if !deleteTodo {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body := r.Body
	defer body.Close()

	bodyJson := helpers.GetJsonFromBody(r)
	bodyMap, ok := bodyJson.(map[string]interface{})
	if !ok {
		http.Error(w, "Failed to decode body", http.StatusInternalServerError)
		return
	}

	if bodyMap["error"] != nil {
		http.Error(w, "Failed to decode body", http.StatusInternalServerError)
	}

	userRequestedTodo := bodyMap["todo"].(string)
	todoExists := mongohelper.FindOneDocument("vite-react-go", "todos", bson.D{{"todo", userRequestedTodo}})
	if todoExists != nil {
		http.Error(w, "Todo already exists", http.StatusBadRequest)
		return
	}

	todo := bson.D{
		{"todo", userRequestedTodo},
	}
	insertedTodo := mongohelper.InsertOneDocument("vite-react-go", "todos", todo)
	if !insertedTodo {
		http.Error(w, "Failed to add todo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ClearTodosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	clearTodos := mongohelper.DeleteManyDocuments("vite-react-go", "todos", bson.D{{}})
	if !clearTodos {
		http.Error(w, "Failed to clear todos", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
