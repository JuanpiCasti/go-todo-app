package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/juanpicasti/go-todo-app/internal/app/handler"
	"github.com/juanpicasti/go-todo-app/internal/app/repository"
	"github.com/juanpicasti/go-todo-app/internal/app/service"
)

func SetupRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	// Initialize repo
	todoRepository := repository.NewTodoRepository(db)
	// Initialize service
	todoService := service.NewTodoService(todoRepository)
	// Initialize handler
	todoHandler := handler.NewTodoHandler(todoService)

	api := router.Group("/api/v1")
	{
		api.GET("/todos", todoHandler.GetAll)
		api.POST("/todos", todoHandler.Create)
	}

	return router
}
