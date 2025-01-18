package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juanpicasti/go-todo-app/app/dtos"
	"github.com/juanpicasti/go-todo-app/app/model"
	"github.com/juanpicasti/go-todo-app/app/repository"
	"github.com/juanpicasti/go-todo-app/config"
)

type AuthService interface {
	Login(req dtos.LoginRequest) (*dtos.LoginResponse, error)
	GenerateToken(user model.UserWithRole) (string, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) *authService {
	return &authService{
		userRepository: repository,
	}
}

func (s *authService) Login(req dtos.LoginRequest) (*dtos.LoginResponse, error) {
	user, err := s.userRepository.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Password != req.Password {
		return nil, errors.New("invalid password")
	}

	token, err := s.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	return &dtos.LoginResponse{Token: token}, nil
}

func (s *authService) GenerateToken(user model.UserWithRole) (string, error) {
	claims := &dtos.Claims{
		UserID:   user.ID,
		RoleID:   user.Role.ID,
		RoleName: user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.CFG.JWTSecret))
}
