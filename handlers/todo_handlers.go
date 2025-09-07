package handlers

import (
	"encoding/json"
	"errors"
	"github.com/benjamint08/vite-react-go-template/helpers"
	"github.com/benjamint08/vite-react-go-template/models"
	"net/http"
	"os"
)

const todoFile = "todos.json"

func readTodos() ([]string, error) {
	if _, err := os.Stat(todoFile); os.IsNotExist(err) {
		if err := os.WriteFile(todoFile, []byte("[]"), 0644); err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile(todoFile)
	if err != nil {
		return nil, err
	}

	var todos []string
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func writeTodos(todos []string) error {
	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, data, 0644)
}

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.ErrorJSON(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	todos, err := readTodos()
	if err != nil {
		helpers.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, todos)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.ErrorJSON(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	var req models.DeleteTodoRequest
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	todos, err := readTodos()
	if err != nil {
		helpers.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	if req.Index < 0 || req.Index >= len(todos) {
		helpers.ErrorJSON(w, http.StatusBadRequest, errors.New("index out of range"))
		return
	}

	todos = append(todos[:req.Index], todos[req.Index+1:]...)

	if err := writeTodos(todos); err != nil {
		helpers.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, nil)
}

func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.ErrorJSON(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	var req models.AddTodoRequest
	if err := helpers.ReadJSON(r, &req); err != nil {
		helpers.ErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	todos, err := readTodos()
	if err != nil {
		helpers.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	for _, todo := range todos {
		if todo == req.Todo {
			helpers.ErrorJSON(w, http.StatusBadRequest, errors.New("todo already exists"))
			return
		}
	}

	todos = append(todos, req.Todo)

	if err := writeTodos(todos); err != nil {
		helpers.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, nil)
}

func ClearTodosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.ErrorJSON(w, http.StatusMethodNotAllowed, errors.New("method not allowed"))
		return
	}

	if err := writeTodos([]string{}); err != nil {
		helpers.ErrorJSON(w, http.StatusInternalServerError, err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, nil)
}
