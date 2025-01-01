package service

import (
	"github.com/juanpicasti/go-todo-app/app/dtos"
	"github.com/juanpicasti/go-todo-app/app/model"
	"github.com/juanpicasti/go-todo-app/app/repository"
)

type TodoService struct {
	repository repository.TodoRepository
}

func NewTodoService(repository repository.TodoRepository) *TodoService {
	return &TodoService{
		repository: repository,
	}
}

func (s *TodoService) GetAll() ([]dtos.TodoResponse, error) {
	todos, err := s.repository.GetAll()
	return mapToResponseDtoSlice(todos), err
}

func (s *TodoService) Create(todoRequest dtos.TodoCreateRequest) (dtos.TodoResponse, error) {
	todo := mapToEntity(todoRequest)
	newTodo, err := s.repository.Create(todo)
	if err != nil {
		return dtos.TodoResponse{}, err
	}
	return mapToResponseDto(newTodo), nil
}

func (s *TodoService) Update(todoRequest dtos.TodoCreateRequest, id int) (dtos.TodoResponse, error) {
	todo := mapToEntity(todoRequest)
	updatedTodo, err := s.repository.Update(todo, id)

	if err != nil {
		return dtos.TodoResponse{}, err
	}

	return mapToResponseDto(updatedTodo), nil
}

func (s *TodoService) GetById(id int) (dtos.TodoResponse, error) {
	todo, err := s.repository.GetById(id)

	if err != nil {
		return dtos.TodoResponse{}, err
	}

	return mapToResponseDto(todo), err
}

func (s *TodoService) Delete(id int) (dtos.TodoResponse, error) {
	todo, err := s.repository.Delete(id)

	if err != nil {
		return dtos.TodoResponse{}, err
	}

	return mapToResponseDto(todo), err
}

func mapToResponseDtoSlice(todos []model.Todo) []dtos.TodoResponse {

	if todos == nil {
		return []dtos.TodoResponse{}
	}

	todoResponses := make([]dtos.TodoResponse, 0, len(todos))

	for _, todo := range todos {
		todoResponses = append(todoResponses, mapToResponseDto(todo))
	}
	return todoResponses
}

func mapToResponseDto(todo model.Todo) dtos.TodoResponse {
	return dtos.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}
}

func mapToEntity(todoRequest dtos.TodoCreateRequest) model.Todo {
	return model.Todo{
		Title:       todoRequest.Title,
		Description: todoRequest.Description,
		Completed:   todoRequest.Completed,
	}
}
