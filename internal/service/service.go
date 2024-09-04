package service

import (
	"bomond-tenis/internal/models"
	"bomond-tenis/internal/repository"
	"bomond-tenis/internal/service/auth"
	"bomond-tenis/internal/service/courts"
	"bomond-tenis/internal/service/users"
)

type Authorization interface {
	SignUp(username, email, password string) (int, error)
	SignIn(email, password string) (string, error)
	Logout(id int) error
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Users interface {
	GetUserById(id int) (models.User, error)
	UpdateUser(id int, user models.User) error
	DeleteUser(id int) error
}

type Courts interface {
	GetCourts() ([]models.Court, error)
	BookCourt(id int) error
	GetCourtById(id int) (models.Court, error)
	DeleteCourt(id int) error
}

type Service struct {
	Authorization
	Users
	Courts
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: auth.NewAuthorizationService(repo),
		Users:         users.NewUserService(repo),
		Courts:        courts.NewCourtService(repo),
	}
}
