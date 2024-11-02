package service

import (
	"crud-api/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const signingKey = "dlkasjdlas5F6SFDSKL3IU2Y3Y298"

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.UserRepositoryInterface
}

func NewAuthService(repo repository.UserRepositoryInterface) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) GenerateToken(login string, password string) (string, error) {
	user, err := a.repo.Find(login, password)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix()},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}
