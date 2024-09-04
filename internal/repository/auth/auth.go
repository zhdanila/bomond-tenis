package auth

import "github.com/jmoiron/sqlx"

type AuthorizationRepository struct {
	db *sqlx.DB
}

func (a AuthorizationRepository) SignUp(username, email, password string) (int, error) {
	panic("implement me")
}

func (a AuthorizationRepository) SignIn(email, password string) (int, error) {
	panic("implement me")
}

func (a AuthorizationRepository) Logout(id int) error {
	panic("implement me")
}

func NewAuthorizationRepository(db *sqlx.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}
