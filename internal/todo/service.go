package todo

import (
	"sync"
)

type Service interface {
	GetTodos() ([]Todo, error)
	CreateTodo(input CreateTodoInput) (Todo, error)
}

type todoService struct {
	todoRepository Repository
}

var instance *todoService
var once sync.Once

func NewTodoService(todoRepository Repository) Service {
	return &todoService{todoRepository}
}

type CreateTodoInput struct {
	Title string `json:"title"`
}

func (s *todoService) GetTodos() ([]Todo, error) {
	return s.todoRepository.GetTodos()
}

func (s *todoService) CreateTodo(input CreateTodoInput) (Todo, error) {
	if input.Title == "" {
		return Todo{}, ErrInvalidTitle
	}

	todo, err := s.todoRepository.CreateTodo(input)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}
