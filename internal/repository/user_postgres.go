package repository

import (
	"standard/internal/model"
	"standard/pkg/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db  *sqlx.DB
	log *logger.Logger
}

func NewUserPostgres(db *sqlx.DB, log *logger.Logger) *UserPostgres {
	return &UserPostgres{
		db:  db,
		log: log,
	}
}

func (r *UserPostgres) Create(request model.UserCreateRequest) (id uuid.UUID, err error) {
	insertQuery := `
	INSERT INTO users (
		full_name,
		email,
		password_hash
	) VALUES ($1, $2, $3) RETURNING id;`

	if err = r.db.Get(&id, insertQuery,
		request.FullName,
		request.Email,
		request.Password,
	); err != nil {
		r.log.Error(err)
		return uuid.Nil, err
	}

	return id, nil
}

func (r *UserPostgres) GetById(request model.UserGetByIdRequest) (user model.User, err error) {
	selectQuery := `
	SELECT 
	    id,
	    full_name,
	    email,
	    created_at,
	    updated_at
	FROM 
	    users
	WHERE 
	    id = $1 
	    AND deleted_at IS NULL;`

	if err = r.db.Get(&user, selectQuery, request.Id); err != nil {
		r.log.Error(err)
		return user, err
	}

	return user, nil
}

func (r *UserPostgres) GetByEmailAndPassword(email string, password string) (user model.User, err error) {
	selectQuery := `
	SELECT 
	    id,
	    full_name,
	    email,
	    created_at,
	    updated_at
	FROM 
	    users
	WHERE 
	    email = $1 
		AND password_hash = $2
	    AND deleted_at IS NULL;`

	if err = r.db.Get(&user, selectQuery, email, password); err != nil {
		r.log.Error(err)
		return user, err
	}

	return user, nil
}
