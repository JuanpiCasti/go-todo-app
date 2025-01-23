package service

import (
	"errors"
	"github.com/juanpicasti/go-todo-app/config"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/juanpicasti/go-todo-app/app/dtos"
	"github.com/juanpicasti/go-todo-app/app/model"
	"github.com/juanpicasti/go-todo-app/app/repository"
)

type AuthService interface {
	Login(req dtos.LoginRequest) (*dtos.LoginResponse, error)
	Register(req dtos.RegisterRequest, roleId int) (*dtos.RegisterResponse, error)
	GenerateToken(user model.UserWithRole) (string, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) AuthService {
	return &authService{
		userRepository: repository,
	}
}

func (s *authService) Login(req dtos.LoginRequest) (*dtos.LoginResponse, error) {
	user, err := s.userRepository.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !checkPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid password")
	}

	token, err := s.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	return &dtos.LoginResponse{Token: token}, nil
}

func (s *authService) Register(req dtos.RegisterRequest, roleId int) (*dtos.RegisterResponse, error) {
	hashedPassword, err := hashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	user := model.User{
		Username: req.Username,
		Password: string(hashedPassword),
		RoleID:   roleId,
	}

	err = s.userRepository.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &dtos.RegisterResponse{Username: user.Username}, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func (s *authService) GenerateToken(user model.UserWithRole) (string, error) {
	claims := &dtos.Claims{
		UserID:   user.ID,
		RoleID:   user.Role.ID,
		RoleName: user.Role.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.CFG.TokenDurationMinutes))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.CFG.JWTSecret))
}
