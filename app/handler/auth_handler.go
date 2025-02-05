package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/juanpicasti/go-todo-app/app/dtos"
	"github.com/juanpicasti/go-todo-app/app/service"
	"github.com/juanpicasti/go-todo-app/app/util"
	"net/http"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginRequest dtos.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerRequest dtos.RegisterRequest
	if errs := util.BindJsonWithErrs(c, &registerRequest); errs != nil {
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	registerResponse, err := h.authService.Register(registerRequest, 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, registerResponse)
}
