package db

import (
	"fmt"
	"sync"

	"github.com/account/app"
	"github.com/goadesign/goa"
)

type AccountModel struct {
	ID       int
	Name     string
	Username string
	Password string
	Email    string
}

type DB struct {
	sync.Mutex
	accounts map[int]*app.Account
}

func NewDB() *DB {
	account := &app.Account{
		ID:       1,
		Name:     "Account 1",
		Username: "Account Username 1",
		Email:    "Account Email 1",
		Password: "Account Password 1",
	}
	return &DB{accounts: map[int]*app.Account{1: account}}
}

// func (seq IDkey) Increment() IDkey {
// 	id = len(UserRepository)
// }

func (db *DB) CreateUser(name string, username string, password string, email string) (*app.Account, error) {
	if name == "internal-error" {
		return nil, goa.ErrInternal("Server internal error")
	}

	account := &app.Account{
		ID:       len(db.accounts) + 1,
		Name:     name,
		Username: username,
		Password: password,
		Email:    email,
	}

	db.accounts[account.ID] = account
	fmt.Println(account.ID)
	return account, nil
}

func (db *DB) GetUser(id int) (*app.Account, error) {

	dbUser := db.accounts[id]

	account := &app.Account{
		ID:       id,
		Name:     dbUser.Name,
		Username: dbUser.Username,
		Password: dbUser.Password,
		Email:    dbUser.Email,
	}
	return account, nil
}

func (db *DB) DeleteUser(id int) error {

	delete(db.accounts, id)
	return nil
}

func (db *DB) UpdateUser(id int, name string, username string, password string, email string) (*app.Account, error) {

	if name == "internal-error" {
		return nil, goa.ErrInternal("Server internal error")
	}

	account := &app.Account{
		ID:       id,
		Name:     name,
		Username: username,
		Password: password,
		Email:    email,
	}

	db.accounts[account.ID] = account
	return account, nil
}
