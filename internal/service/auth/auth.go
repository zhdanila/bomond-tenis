package auth

import (
	"bomond-tenis/internal/repository"
	"bomond-tenis/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

const signingKey = "*#jcpk040(OJF$OJOI)#Ujf094*"

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (a AuthorizationService) GenerateToken(username, password string) (string, error) {
	user, err := a.repo.GetUser(username, utils.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(user.Id),
		"exp": time.Now().Add(time.Hour * 6).Unix(),
	})

	return token.SignedString([]byte(signingKey))
}

func (a AuthorizationService) ParseToken(token string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorizationService) SignUp(username, email, password string) (int, error) {
	password = utils.GeneratePasswordHash(password)

	return a.repo.SignUp(username, email, password)
}

func (a AuthorizationService) SignIn(email, password string) (int, error) {
	panic("implement me")
}

func (a AuthorizationService) Logout(id int) error {
	panic("implement me")
}
