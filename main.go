package main

import (
	"encoding/json"
	"net/http"
)

type TodoItem struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	IsComplete  bool   `json:"is_complete"`
}

type todoHandlers struct {
	store map[int]TodoItem
}

func (h *todoHandlers) get(writer http.ResponseWriter, request *http.Request) {
	todos := make([]TodoItem, len(h.store))

	i := 0
	for _, todo := range h.store {
		todos[i] = todo
		i++
	}

	bytes, err := json.Marshal(todos)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
	}

	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}

func newTodoHandlers() *todoHandlers {
	return &todoHandlers{
		store: map[int]TodoItem{
			1: {
				ID:          1,
				Description: "Buy milk",
				IsComplete:  false,
			},
		},
	}
}

func main() {
	todoHandlers := newTodoHandlers()

	http.HandleFunc("/todos", todoHandlers.get)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		panic(err)
	}
}
