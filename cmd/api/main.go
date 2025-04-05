package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/igor-marchi/to-do/internal/config"
	"github.com/igor-marchi/to-do/internal/db"
	"github.com/igor-marchi/to-do/internal/todo"
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

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("API rodando em http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
