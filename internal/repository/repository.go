package repository

import (
	"standard/internal/model"
	"standard/pkg/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserI
}

func NewRepository(db *sqlx.DB, log *logger.Logger) *Repository {
	return &Repository{
		UserI: NewUserPostgres(db, log),
	}
}

type UserI interface {
	Create(request model.UserCreateRequest) (id uuid.UUID, err error)
	GetById(request model.UserGetByIdRequest) (user model.User, err error)
	GetByEmailAndPassword(email string, password string) (user model.User, err error)
}
