package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID  `json:"id" db:"id"`
	FullName  string     `json:"fullName" db:"full_name"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"password" db:"-"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

type UserCreateRequest struct {
	FullName string `json:"fullName" db:"full_name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`
}

type UserGetByIdRequest struct {
	Id uuid.UUID
}
