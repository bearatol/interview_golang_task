package service

import (
	"fmt"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/config"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	conf         *config.Config
	repoUser     RepoUser
	repoProducts RepoProducts
	repoPrices   RepoPrices
	priceGen     PriceGenerator
	authGen      AuthGenerator
}

func NewService(
	conf *config.Config,
	repo *repository.Repository,
	priceGen PriceGenerator,
	authGen AuthGenerator,
) *Service {
	return &Service{
		conf:         conf,
		repoUser:     repo,
		repoProducts: repo,
		repoPrices:   repo,
		priceGen:     priceGen,
		authGen:      authGen,
	}
}

func (s *Service) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *Service) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *Service) extractJWTClaims(tokenStr, key string) (claims jwt.MapClaims, err error) {
	hmacSecretString := key
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid jwt token")
}

func (s *Service) login(token string) (string, error) {
	res, err := s.extractJWTClaims(token, s.conf.JWTKey)
	if err != nil {
		return "", err
	}
	loginJWT, exist := res["login"]
	if !exist {
		return "", fmt.Errorf("invalid login")
	}
	login, ok := loginJWT.(string)
	if !ok {
		return "", fmt.Errorf("invalid login")
	}
	return login, nil
}
