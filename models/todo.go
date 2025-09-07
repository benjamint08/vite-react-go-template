package models

type AddTodoRequest struct {
	Todo string `json:"todo"`
}

type DeleteTodoRequest struct {
	Index int `json:"index"`
}
