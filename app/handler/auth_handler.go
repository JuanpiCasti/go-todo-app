package handler

import (
	"github.com/juanpicasti/go-todo-app/app/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanpicasti/go-todo-app/app/dtos"
	"github.com/juanpicasti/go-todo-app/app/service"
)

type AuthHandler struct {
	authService       service.AuthService
	passwordValidator util.Validator
}

func NewAuthHandler(authService service.AuthService, validator util.Validator) *AuthHandler {
	return &AuthHandler{authService, validator}
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
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.passwordValidator.Validate(registerRequest.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registerResponse, err := h.authService.Register(registerRequest, 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, registerResponse)
}
