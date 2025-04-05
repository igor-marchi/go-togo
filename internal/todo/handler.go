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

// GetTodos godoc
// @Summary      Lista todas as tarefas
// @Description  Retorna uma lista de todos os todos
// @Tags         todos
// @Produce      json
// @Success      200  {array}  todo.Todo
// @Router       /todos [get]

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

// CreateTodo godoc
// @Summary      Cria uma nova tarefa
// @Description  Cria uma tarefa com título
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body      todo.CreateTodoInput  true  "Dados da tarefa"
// @Success      201   {object}  todo.Todo
// @Failure      400   {string}  string  "Erro de validação"
// @Router       /todos [post]

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
