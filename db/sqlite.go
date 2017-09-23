package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/account/app"

	"github.com/goadesign/goa"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	*sql.DB
}

func NewSqliteDB() *SqliteDB {
	db, err := sql.Open("sqlite3", "foo.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	 CREATE TABLE IF NOT EXISTS accounts (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, name text, username text, password text, email text)
	 `
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}

	return &SqliteDB{db}
}

func (db *SqliteDB) CreateUser(name string, username string, password string, email string) (*app.Account, error) {
	fmt.Println(db)

	tx, err := db.Begin()
	if err != nil {
		return nil, goa.ErrInternal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO accounts(name, username, password, email) values (?, ?, ?, ?)")
	if err != nil {
		return nil, goa.ErrInternal(err)
	}

	account, err := stmt.Exec(name, username, password, email)
	if err != nil {
		return nil, goa.ErrInternal(err)
	}

	tx.Commit()

	fmt.Print(account)

	res := &app.Account{
		Name:     name,
		Username: username,
		Password: password,
		Email:    email,
		ID:       1,
	}

	fmt.Print("here")
	fmt.Println(res)

	return res, nil
}

func (db *SqliteDB) GetUser(id int) (*app.Account, error) {
	stmt, err := db.Prepare("SELECT * FROM accounts WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var name, username, password, email string
	err = stmt.QueryRow(id).Scan(&id, &name, &username, &password, &email)
	if err != nil {
		return nil, goa.ErrNotFound(err)
	}

	fmt.Println("Results: ", id, name, username, password, email)
	res := &app.Account{
		ID:       id,
		Name:     name,
		Username: username,
		Password: password,
		Email:    email,
	}

	return res, nil
}

func (db *SqliteDB) DeleteUser(id int) error {

	stmt, err := db.Prepare("DELETE FROM accounts WHERE id = ?")
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return goa.ErrNotFound(err)
	}

	return nil

}

func (db *SqliteDB) UpdateUser(id int, name string, username string, password string, email string) (*app.Account, error) {

	tx, err := db.Begin()
	if err != nil {
		return nil, goa.ErrInternal(err)
	}

	stmt, err := db.Prepare("UPDATE accounts set name=?, username=?, password=?, email=? WHERE id = ?")
	defer stmt.Close()

	// fmt.Println(id)
	// fmt.Println(name)
	// fmt.Println(username)
	// fmt.Println(password)
	// fmt.Println(email)

	_, err = stmt.Exec(name, username, password, email, id)

	if err != nil {
		return nil, goa.ErrNotFound(err)
	}

	tx.Commit()

	res := &app.Account{
		Name:     name,
		Username: username,
		Password: password,
		Email:    email,
		ID:       id,
	}

	return res, nil

}
