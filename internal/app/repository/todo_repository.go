package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/juanpicasti/go-todo-app/internal/app/model"
)

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	Create(todo model.Todo) (model.Todo, error)
}

type todoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *todoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	err := r.db.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) Create(todo model.Todo) (model.Todo, error) {
	var newTodo model.Todo
	query := `
        INSERT INTO todos (title, description) 
        VALUES ($1, $2) 
        RETURNING id, title, description, completed
    `
	err := r.db.
		QueryRow(
			query,
			todo.Title,
			todo.Description).
		Scan(
			&newTodo.ID,
			&newTodo.Title,
			&newTodo.Description,
			&newTodo.Completed)

	if err != nil {
		return model.Todo{}, err
	}
	return newTodo, nil
}
