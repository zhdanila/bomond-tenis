package users

import (
	"bomond-tenis/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) GetUserById(id int) (models.User, error) {
	panic("implement me")
}

func (u UserRepository) UpdateUser(id int, user models.User) error {
	panic("implement me")
}

func (u UserRepository) DeleteUser(id int) error {
	panic("implement me")
}
