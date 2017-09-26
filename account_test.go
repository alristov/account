package main

import (
	"context"
	"testing"

	"github.com/account/app"
	"github.com/account/app/test"
	"github.com/account/db"

	"github.com/goadesign/goa"
)

var payloadCreate = &app.CreateUserAccountPayload{
	Email:    "EmailTest@test.com",
	Name:     "NameTest",
	Username: "UsernameTest",
	Password: "PasswordTest",
}

var payloadBadRequest = &app.CreateUserAccountPayload{
	Email:    "EmailTest@test.com",
	Name:     "NameTest",
	Username: "UsernameTest",
	Password: "PasswordTest",
}

var payloadInternalErr = &app.CreateUserAccountPayload{
	Email:    "EmailTest@test.com",
	Name:     "NameTest",
	Username: "UsernameTest",
	Password: "PasswordTest",
}

var payloadUpdate = &app.UpdateUserAccountPayload{
	Email:    "EmailTest@test.com",
	Name:     "NameTest",
	Username: "UsernameTest",
	Password: "PasswordTest",
}

var payloadBadRequestUpdate = &app.UpdateUserAccountPayload{
	Username: "UsernameTest",
	Password: "PasswordTest",
}

var payloadNotFound = &app.UpdateUserAccountPayload{
	Email:    "EmailTest@test.com",
	Name:     "NameTest",
	Username: "UsernameTest",
	Password: "PasswordTest",
}

var payloadOK = &app.UpdateUserAccountPayload{
	Email:    "EmailTest@test.com",
	Name:     "NameTest",
	Username: "UsernameTest",
	Password: "PasswordTest",
}

var payloadInternalErrUpdate = &app.UpdateUserAccountPayload{
	Email:    "EmailTest@test.com",
	Name:     "NameTest",
	Username: "UsernameTest",
	Password: "PasswordTest",
}

var payloadInternalErrCreate = &app.CreateUserAccountPayload{
	Email:    "EmailTest@test.com",
	Name:     "internal-error",
	Username: "UsernameTest",
	Password: "PasswordTest",
}

var (
	service  = goa.New("account-test")
	database = db.NewDB()
	ctrl     = NewAccountController(service, database)
)

func TestCreateUserAccountCreated(t *testing.T) {

	_, account := test.CreateUserAccountCreated(t, context.Background(), service, ctrl, payloadCreate)

	if account == nil {
		t.Fatal("Nil account!!!")
	}

}

func TestCreateUserAccountInternalServerError(t *testing.T) {
	_, err := test.CreateUserAccountInternalServerError(t, context.Background(), service, ctrl, payloadInternalErr)

	if err == nil {
		t.Fatal("Nil error")
	}

}

func TestCreateUserAccountBadRequest(t *testing.T) {
	_, err := test.CreateUserAccountBadRequest(t, context.Background(), service, ctrl, payloadBadRequest)

	if err == nil {
		t.Fatal("Nil error")
	}

}

func TestGetUserAccountOK(t *testing.T) {
	_, account := test.GetUserAccountOK(t, context.Background(), service, ctrl, 1)

	if account == nil {
		t.Fatal("Error, not OK")
	}
}

func TestGetUserAccountNotFound(t *testing.T) {
	test.GetUserAccountNotFound(t, context.Background(), service, ctrl, 2)
}

func TestUpdateUserAccountOK(t *testing.T) {
	_, account := test.UpdateUserAccountOK(t, context.Background(), service, ctrl, 4, payloadOK)

	if account == nil {
		t.Fatal("Nil error")
	}
}

func TestUpdateUserAccountNotFound(t *testing.T) {
	test.UpdateUserAccountNotFound(t, context.Background(), service, ctrl, 2, payloadUpdate)
}

func TestUpdateUserAccountBadRequest(t *testing.T) {
	_, err := test.UpdateUserAccountBadRequest(t, context.Background(), service, ctrl, 1, payloadBadRequestUpdate)

	if err == nil {
		t.Fatal("Error nil")
	}

}

func TestUpdateUserAccountInternalServerError(t *testing.T) {
	_, err := test.UpdateUserAccountInternalServerError(t, context.Background(), service, ctrl, 3, payloadInternalErrUpdate)

	if err == nil {
		t.Fatal("Nil error")
	}

}

func TestDeleteUserAccountNotFound(t *testing.T) {
	test.DeleteUserAccountNotFound(t, context.Background(), service, ctrl, 2)
}

func TestDeleteUserAccountBadRequest(t *testing.T) {
	_, err := test.DeleteUserAccountBadRequest(t, context.Background(), service, ctrl, 5)

	if err == nil {
		t.Fatal("Error, can't delete")
	}
}

func TestDeleteUserAccountInternalServerError(t *testing.T) {
	test.DeleteUserAccountInternalServerError(t, context.Background(), service, ctrl, 3)
}

func TestDeleteUserAccountNoContent(t *testing.T) {
	test.DeleteUserAccountNoContent(t, context.Background(), service, ctrl, 1)
}
