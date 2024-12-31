package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/juanpicasti/go-todo-app/internal/app/dtos"
	"github.com/juanpicasti/go-todo-app/internal/app/model"
	"github.com/juanpicasti/go-todo-app/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestTodoService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockTodoRepository(ctrl)
	service := NewTodoService(mockRepo)

	t.Run("Success - Returns todos", func(t *testing.T) {
		// Arrange
		mockTodos := []model.Todo{
			{ID: 1, Title: "Test1", Description: "Desc1", Completed: false},
			{ID: 2, Title: "Test2", Description: "Desc2", Completed: true},
		}
		expectedResponse := []dtos.TodoResponse{
			{ID: 1, Title: "Test1", Description: "Desc1", Completed: false},
			{ID: 2, Title: "Test2", Description: "Desc2", Completed: true},
		}
		mockRepo.EXPECT().GetAll().Return(mockTodos, nil)

		// Act
		result, err := service.GetAll()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, result)
	})
}
