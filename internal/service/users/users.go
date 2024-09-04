package users

import (
	"bomond-tenis/internal/models"
	"bomond-tenis/internal/repository"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUserById(id int) (models.User, error) {
	panic("")
}

func (u *UserService) UpdateUser(id int, user models.User) error {
	panic("")
}

func (u *UserService) DeleteUser(id int) error {
	panic("")
}
