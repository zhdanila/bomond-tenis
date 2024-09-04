package auth

import (
	"bomond-tenis/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const UsersTable = "users"

type AuthorizationRepository struct {
	db *sqlx.DB
}

func NewAuthorizationRepository(db *sqlx.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

func (a AuthorizationRepository) GetUser(email, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", UsersTable)
	err := a.db.Get(&user, query, email, password)

	return user, err
}

func (a AuthorizationRepository) SignUp(username, email, password string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password) values ($1, $2, $3) RETURNING id", UsersTable)

	row := a.db.QueryRow(query, username, email, password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a AuthorizationRepository) SignIn(email, password string) (int, error) {
	panic("implement me")
}

func (a AuthorizationRepository) Logout(id int) error {
	panic("implement me")
}
