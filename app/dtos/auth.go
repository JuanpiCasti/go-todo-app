package dtos

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username       string `json:"username" binding:"required,min=5"`
	Password       string `json:"password" binding:"required,min=8"`
	PasswordRepeat string `json:"password_repeat" binding:"required,eqfield=Password"`
	PhoneNumber    string `json:"phone_number" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Name           string `json:"name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
}

type RegisterResponse struct {
	Username string `json:"username"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type Claims struct {
	UserID   int    `json:"user_id"`
	RoleID   int    `json:"role_id"`
	RoleName string `json:"role_name"`
	jwt.RegisteredClaims
}
