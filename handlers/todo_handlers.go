package handlers

import (
	"encoding/json"
	"github.com/benjamint08/vite-react-go-template/helpers"
	"net/http"
	"os"
)

func checkTodosFile() {
	_, err := os.Stat("todos.json")
	if os.IsNotExist(err) {
		err := os.WriteFile("todos.json", []byte("[]"), 0644)
		if err != nil {
			panic("Failed to create todos")
		}
	}
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	checkTodosFile()
	data, err := os.ReadFile("todos.json")
	if err != nil {
		http.Error(w, "Failed to read todos", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body := r.Body
	defer body.Close()
	checkTodosFile()
	data, err := os.ReadFile("todos.json")
	if err != nil {
		http.Error(w, "Failed to read todos", http.StatusInternalServerError)
		return
	}

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

	var currentTodos []string
	err = json.Unmarshal(data, &currentTodos)
	if err != nil {
		http.Error(w, "Failed to decode todos", http.StatusInternalServerError)
		return
	}
	index := int(bodyMap["index"].(float64))
	currentTodos = append(currentTodos[:index], currentTodos[index+1:]...)
	newData, err := json.Marshal(currentTodos)
	if err != nil {
		http.Error(w, "Failed to encode todos", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("todos.json", newData, 0644)
	if err != nil {
		http.Error(w, "Failed to write todos", http.StatusInternalServerError)
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
	checkTodosFile()
	data, err := os.ReadFile("todos.json")
	if err != nil {
		http.Error(w, "Failed to read todos", http.StatusInternalServerError)
		return
	}

	bodyJson := helpers.GetJsonFromBody(r)
	bodyMap, ok := bodyJson.(map[string]interface{})
	if !ok {
		http.Error(w, "Failed to decode body", http.StatusInternalServerError)
		return
	}

	if bodyMap["error"] != nil {
		http.Error(w, "Failed to decode body", http.StatusInternalServerError)
	}

	var currentTodos []string
	err = json.Unmarshal(data, &currentTodos)
	if err != nil {
		http.Error(w, "Failed to decode todos", http.StatusInternalServerError)
		return
	}
	for _, todo := range currentTodos {
		if todo == bodyMap["todo"].(string) {
			http.Error(w, "Todo already exists", http.StatusBadRequest)
			return
		}
	}
	currentTodos = append(currentTodos, bodyMap["todo"].(string))
	newData, err := json.Marshal(currentTodos)
	if err != nil {
		http.Error(w, "Failed to encode todos", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("todos.json", newData, 0644)
	if err != nil {
		http.Error(w, "Failed to write todos", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ClearTodosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	checkTodosFile()
	err := os.WriteFile("todos.json", []byte("[]"), 0644)
	if err != nil {
		http.Error(w, "Failed to clear todos", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
