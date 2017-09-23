package db

import (
	"github.com/account/app"
)

type Repository interface {
	CreateUser(name string, username string, password string, email string) (*app.Account, error)
	GetUser(id int) (*app.Account, error)
	DeleteUser(id int) error
	UpdateUser(id int, name string, username string, password string, email string) (*app.Account, error)
}
