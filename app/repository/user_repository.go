package repository

//go:generate go run github.com/golang/mock/mockgen -destination=mock_user_repository.go -package=repository github.com/juanpicasti/go-todo-app/internal/app/repository UserRepository

import (
	"github.com/juanpicasti/go-todo-app/app/model"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	FindByUsername(string) (*model.UserWithRole, error)
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByUsername(username string) (*model.UserWithRole, error) {
	query := `
    SELECT
      u.*,
      r.id AS "role.id",
      r.name AS "role.name"
    FROM todo_app.app_users u
    JOIN todo_app.app_roles r ON u.role_id = r.id
    WHERE u.username = $1
  `
	userWithRole := model.UserWithRole{}
	err := r.db.Get(&userWithRole, query, username)
	if err != nil {
		return nil, err
	}

	return &userWithRole, nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	query := `
	INSERT INTO todo_app.app_users (username, password, role_id, active)
	VALUES ($1, $2, $3, true)
  `
	_, err := r.db.Exec(query, user.Username, user.Password, user.RoleID)
	return err
}
