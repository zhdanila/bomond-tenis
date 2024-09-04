package auth

import "bomond-tenis/internal/repository"

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (a AuthorizationService) SignUp(username, email, password string) (int, error) {
	panic("implement me")
}

func (a AuthorizationService) SignIn(email, password string) (int, error) {
	panic("implement me")
}

func (a AuthorizationService) Logout(id int) error {
	panic("implement me")
}
