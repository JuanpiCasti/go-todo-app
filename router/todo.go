package router

import (
	"github.com/juanpicasti/go-todo-app/app/handler"
	"github.com/juanpicasti/go-todo-app/app/middleware"
	"github.com/juanpicasti/go-todo-app/app/repository"
	"github.com/juanpicasti/go-todo-app/app/service"
)

func (r *Router) initTodoHandler() {
	todoRepo := repository.NewTodoRepository(r.db)
	todoService := service.NewTodoService(todoRepo)
	r.todoHandler = handler.NewTodoHandler(todoService)
}

func (r *Router) setupApiRoutes() {
	api := r.engine.Group("/api/v1")

	api.Use(middleware.AuthMiddleware())
	api.Use(middleware.RoleMiddleware(map[string]bool{
		"admin": true,
		"user":  true,
	}))

	api.GET("/todos", r.todoHandler.GetAll)
	api.GET("/todos/:id", r.todoHandler.GetById)
	api.POST("/todos", r.todoHandler.Create)
	api.PUT("/todos/:id", r.todoHandler.Update)
	api.DELETE("/todos/:id", r.todoHandler.Delete)
}
