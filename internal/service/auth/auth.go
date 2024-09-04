package auth

import (
	"bomond-tenis/internal/repository"
	"bomond-tenis/pkg/utils"
	"github.com/go-openapi/errors"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"strings"
	"time"
)

const (
	signingKey = "*#jcpk040(OJF$OJOI)#Ujf094*"
	tokenTTL   = time.Hour * 6
)

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (a AuthorizationService) GenerateToken(email, password string) (string, error) {
	user, err := a.repo.GetUser(email, utils.GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(user.Id),
		"exp": time.Now().Add(tokenTTL).Unix(),
	})

	return token.SignedString([]byte(signingKey))
}

func (a AuthorizationService) ParseToken(header string) (int, error) {
	if header == "" {
		return 0, errors.New(500, "Empty header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return 0, errors.New(500, "Invalid header")
	}

	claims, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(500, "unexpected signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	id, err := claims.Claims.GetSubject()
	if err != nil {
		return 0, err
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return intId, nil
}

func (a AuthorizationService) SignUp(username, email, password string) (int, error) {
	password = utils.GeneratePasswordHash(password)

	return a.repo.SignUp(username, email, password)
}

func (a AuthorizationService) SignIn(email, password string) (string, error) {
	return a.GenerateToken(email, password)
}

func (a AuthorizationService) Logout(id int) error {
	panic("implement me")
}
