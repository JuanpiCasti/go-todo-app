package dtos

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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
