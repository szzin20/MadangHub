package entity

import (
	"time"
)

type Main struct {
	Id        string
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Phone     string
	Address   string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserDataInterface interface {
	Register(data Main) (row int, err error)
	Login(email, username, password string) (Main, string, error)
	// GetData(ID uuid.UUID) (Main, error)
}

type UseCaseInterface interface {
	Register(data Main) (row int, err error)
	Login(email, username, password string) (Main, string, error)
}