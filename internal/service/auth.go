package service

import (
	"errors"
	"test_task/internal/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	signingKey = "qrkjk#4#35FSFJ"
	tokenTTL   = 2 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	Phone string `json:"phone"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignIn(phone string) error {
	user, err := s.repo.GetUser(phone)
	if err != nil {
		return err
	}

	if user.ID == 0 {
		err = s.repo.CreateUser(phone)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *AuthService) GenerateToken(phone string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		phone,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of types *token.Claims")
	}
	return claims.Phone, nil
}
