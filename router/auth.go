package router

import (
	"github.com/juanpicasti/go-todo-app/app/handler"
	"github.com/juanpicasti/go-todo-app/app/repository"
	"github.com/juanpicasti/go-todo-app/app/service"
	"github.com/juanpicasti/go-todo-app/app/util"
)

func (r *Router) initAuthHandler() {
	userRepo := repository.NewUserRepository(r.db)
	authService := service.NewAuthService(userRepo)
	r.authHandler = handler.NewAuthHandler(authService, util.NewPasswordValidator())
}

func (r *Router) setupAuthRoutes() {
	r.engine.POST("/login", r.authHandler.Login)
	r.engine.POST("/register", r.authHandler.Register)
}
