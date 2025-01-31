package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/juanpicasti/go-todo-app/app/handler"
	"github.com/juanpicasti/go-todo-app/app/middleware"
	"github.com/juanpicasti/go-todo-app/config"
	"github.com/rs/zerolog/log"
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
	router.setTrustedProxies()
	router.initializeHandlers()
	router.setupMiddleware()
	router.setupRoutes()
	return router.engine
}

func (r *Router) setupRoutes() {
	r.setupAuthRoutes()
	r.setupApiRoutes()
}

func (r *Router) initializeHandlers() {
	r.initAuthHandler()
	r.initTodoHandler()
}

func (r *Router) setTrustedProxies() {
	err := r.engine.SetTrustedProxies(config.CFG.TrustedProxies)
	if err != nil {
		log.Error().Err(err).Msg("Could not set trusted proxies")
		panic(err)
	}
}
