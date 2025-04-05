package todo

import "database/sql"

type Repository interface {
	GetTodos() ([]Todo, error)
	CreateTodo(input CreateTodoInput) (Todo, error)
}

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) Repository {
	return &todoRepository{db}
}

func (r *todoRepository) GetTodos() ([]Todo, error) {
	rows, err := r.db.Query("SELECT id, title, done FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Done); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *todoRepository) CreateTodo(input CreateTodoInput) (Todo, error) {
	result, err := r.db.Exec("INSERT INTO todos (title, done) VALUES (?, ?)", input.Title, false)
	if err != nil {
		return Todo{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Todo{}, err
	}

	todo := Todo{
		ID:    int(id),
		Title: input.Title,
		Done:  false,
	}

	return todo, nil
}
