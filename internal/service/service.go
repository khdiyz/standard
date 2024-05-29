package service

import (
	"standard/internal/model"
	"standard/internal/repository"
	"standard/pkg/logger"

	"github.com/google/uuid"
)

type Service struct {
	UserI
	Authorization
}

func NewService(repos *repository.Repository, log logger.Logger) *Service {
	return &Service{
		UserI:         NewUserService(repos, log),
		Authorization: NewAuthService(repos, log),
	}
}

type Authorization interface {
	GenerateToken(email string, password string) (token string, err error)
	ParseToken(accessToken string) (uuid.UUID, error)
}

type UserI interface {
	CreateUser(user model.UserCreateRequest) (id uuid.UUID, err error)
}
