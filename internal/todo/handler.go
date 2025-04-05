package todo

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {

	todos, err := h.service.GetTodos()
	if err != nil {
		http.Error(w, "Error fetching todos", http.StatusInternalServerError)
		return
	}

	if len(todos) == 0 {
		todos = []Todo{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (h *Handler) CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var input CreateTodoInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	todo, err := h.service.CreateTodo(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
