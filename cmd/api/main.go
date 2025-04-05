package main

// @title           Todo API
// @version         1.0
// @description     Uma API simples de tarefas feita em Go com SQLite.

// @contact.name   Igor Marchi

// @host      localhost:8080
// @BasePath  /todos

import (
	"fmt"
	"log"
	"net/http"

	"github.com/igor-marchi/to-do/internal/config"
	"github.com/igor-marchi/to-do/internal/db"
	"github.com/igor-marchi/to-do/internal/todo"

	_ "github.com/igor-marchi/go-todo/docs" // importa os arquivos gerados
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	cfg := config.Load()

	sqliteDB := db.Connect(cfg.DBPath)
	repo := todo.NewTodoRepository(sqliteDB)
	service := todo.NewTodoService(repo)
	handler := todo.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetTodos(w, r)
		case http.MethodPost:
			handler.CreateTodoHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	mux.HandleFunc("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("API rodando em http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
