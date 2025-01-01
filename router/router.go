package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juanpicasti/go-todo-app/app/handler"
	"github.com/juanpicasti/go-todo-app/app/repository"
	"github.com/juanpicasti/go-todo-app/app/service"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Initialize repo
	todoRepository := repository.NewTodoRepository()
	// Initialize service
	todoService := service.NewTodoService(todoRepository)
	// Initialize handler
	todoHandler := handler.NewTodoHandler(todoService)

	api := router.Group("/api/v1")
	{
		api.GET("/todos", todoHandler.GetAll)
		api.GET("/todos/:id", todoHandler.GetById)
		api.POST("/todos", todoHandler.Create)
		api.PUT("/todos/:id", todoHandler.Update)
		api.DELETE("/todos/:id", todoHandler.Delete)
	}

	return router
}
