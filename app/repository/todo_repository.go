package repository

//go:generate go run github.com/golang/mock/mockgen -destination=mock_todo_repository.go -package=repository github.com/juanpicasti/go-todo-app/internal/app/repository TodoRepository

import (
	"github.com/juanpicasti/go-todo-app/app/model"
	"github.com/juanpicasti/go-todo-app/database"
	"log"

	"github.com/jmoiron/sqlx"
)

type TodoRepository interface {
	GetAll() ([]model.Todo, error)
	Create(todo model.Todo) (model.Todo, error)
	Update(todo model.Todo, id int) (model.Todo, error)
	GetById(id int) (model.Todo, error)
	Delete(id int) (model.Todo, error)
}

type todoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository() *todoRepository {
	return &todoRepository{database.DB}
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

func (r *todoRepository) Update(todo model.Todo, id int) (model.Todo, error) {
	var updatedTodo model.Todo
	query := `
		UPDATE todos
		SET title = $1, description = $2, completed = $3
		WHERE id = $4
		RETURNING id, title, description, completed
	`

	err := r.db.
		QueryRow(
			query,
			todo.Title,
			todo.Description,
			todo.Completed,
			id).
		Scan(
			&updatedTodo.ID,
			&updatedTodo.Title,
			&updatedTodo.Description,
			&updatedTodo.Completed)

	if err != nil {
		return model.Todo{}, err
	}

	return updatedTodo, nil
}

func (r *todoRepository) GetById(id int) (model.Todo, error) {
	var todo model.Todo
	err := r.db.Get(&todo, "SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		return model.Todo{}, err
	}
	return todo, err
}

func (r *todoRepository) Delete(id int) (model.Todo, error) {
	todo, err := r.GetById(id)

	if err != nil {
		return model.Todo{}, err
		// Return custom error
	}

	_, err = r.db.Exec("DELETE FROM todos WHERE id=$1", id)

	if err != nil {
		log.Println(err.Error())
		return model.Todo{}, err
	}

	return todo, nil
}