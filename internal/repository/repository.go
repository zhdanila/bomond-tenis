package repository

import (
	"bomond-tenis/internal/models"
	"bomond-tenis/internal/repository/auth"
	"bomond-tenis/internal/repository/courts"
	"bomond-tenis/internal/repository/users"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	SignUp(username, email, password string) (int, error)
	SignIn(email, password string) (int, error)
	Logout(id int) error
}

type Users interface {
	GetUser(id int) (models.User, error)
	UpdateUser(id int, user models.User) error
	DeleteUser(id int) error
}

type Courts interface {
	GetCourts() ([]models.Court, error)
	BookCourt(id int) error
	GetCourtById(id int) (models.Court, error)
	DeleteCourt(id int) error
}

type Repository struct {
	Authorization
	Users
	Courts
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: auth.NewAuthorizationRepository(db),
		Users:         users.NewUserRepository(db),
		Courts:        courts.NewCourtsRepository(db),
	}
}
