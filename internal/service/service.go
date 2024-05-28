package service

import (
	"standard/internal/repository"
	"standard/pkg/logger"
)

type Service struct {
}

func NewService(repos *repository.Repository, log *logger.Logger) *Service {
	return &Service{}
}
