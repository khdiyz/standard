package service

import (
	"standard/internal/model"
	"standard/internal/repository"
	"standard/pkg/helper"
	"standard/pkg/logger"

	"github.com/google/uuid"
)

type UserService struct {
	repo repository.UserI
	log  logger.Logger
}

func NewUserService(repo repository.UserI, log logger.Logger) *UserService {
	return &UserService{
		repo: repo,
		log:  log,
	}
}

func (s *UserService) CreateUser(user model.UserCreateRequest) (id uuid.UUID, err error) {
	user.Password, err = helper.GenerateHash(user.Password)
	if err != nil {
		return uuid.Nil, err
	}

	return s.repo.Create(user)
}
