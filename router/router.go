package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/juanpicasti/go-todo-app/app/handler"
	"github.com/juanpicasti/go-todo-app/app/middleware"
)

type Router struct {
	engine      *gin.Engine
	db          *sqlx.DB
	todoHandler *handler.TodoHandler
	authHandler *handler.AuthHandler
}

func NewRouter(db *sqlx.DB) *Router {
	return &Router{
		engine: gin.New(),
		db:     db,
	}
}

func (r *Router) setupMiddleware() {
	r.engine.Use(middleware.LoggerMiddleware())
	r.engine.Use(gin.Recovery())
	r.engine.Use(middleware.GetCorsMiddleware())
}

func SetupRouter(db *sqlx.DB) *gin.Engine {
	router := NewRouter(db)
	router.initializeHandlers()
	router.setupMiddleware()
	router.setupRoutes()
	return router.engine
}

func (r *Router) setupRoutes() {
	r.setupAuthRoutes()

	// API routes
	api := r.engine.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())
	api.Use(middleware.RoleMiddleware(map[string]bool{
		"admin": true,
		"user":  true,
	}))

	{
		r.setupTodoRoutes(api)
	}
}

func (r *Router) initializeHandlers() {
	r.initAuthHandler()
	r.initTodoHandler()
}
