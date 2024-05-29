package service

import (
	"errors"
	"standard/internal/config"
	"standard/internal/repository"
	"standard/pkg/helper"
	"standard/pkg/logger"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type AuthService struct {
	repo repository.UserI
	log  logger.Logger
}

func NewAuthService(repo repository.UserI, log logger.Logger) *AuthService {
	return &AuthService{
		repo: repo,
		log:  log,
	}
}

type jwtCustomClaim struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"userId"`
}

func (s *AuthService) GenerateToken(email string, password string) (token string, err error) {
	hash, err := helper.GenerateHash(password)
	if err != nil {
		return "", err
	}

	user, err := s.repo.GetByEmailAndPassword(email, hash)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtCustomClaim{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(config.GetConfig().JWTExpired)).Unix(),
			Issuer:    config.GetConfig().JWTIssuer,
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return jwtToken.SignedString([]byte(config.GetConfig().JWTSecret))
}

func (s *AuthService) ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*jwtCustomClaim)
	if !ok {
		return uuid.Nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}
