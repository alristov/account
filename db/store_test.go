package db

import "testing"

// func TestCreateUser(
//   service := goa.New("User testing")
//   db := dstore.
// )

func TestCreateUser(t *testing.T) {
	var db Repository = NewDB()
	account := db.CreateUser("Some name", "Some username", "Some password", "Some email")
	if account.Username != "Some username" {
		t.Fatal("Username not setup correctly")
	}
}

func TestGetUser(t *testing.T) {
	var db Repository = NewDB()
	account := db.GetUser(1)
	if account.ID != 1 {
		t.Fatal("You should get username with ID : 1 ")
	}

}

func TestUpdateUser(t *testing.T) {
	var db Repository = NewDB()
	account := db.UpdateUser(1, "Name", "Username", "UserPassword", "UserEmail")

	if account.id != 1 {
		t.Fatal("ID should be 1")
	}

}

func TestDeleteUser(t *testing.T) {
	var db Repository = NewDB()

	db.DeleteUser(1)

}
