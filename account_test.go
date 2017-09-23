package main

import (
	"context"
	"testing"

	"account/app"
	"account/app/test"
	"account/db"

	"github.com/goadesign/goa"
)

var payload = &app.CreateUserAccountPayload{
	Email:    "EmailTest@test.com",
	Name:     "NameTest",
	Username: "UsernameTest",
	Password: "PasswordTest",
}

var payloadInternalErr = &app.CreateUserAccountPayload {
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

	_, account := test.CreateUserAccountCreated(t, context.Background(), service, ctrl, payload)

	if account == nil {
		t.Fatal("Nil account!!!")
	}


}

func TestCreateUserAccountInternalServerError(t *testing.T){
	_, _ account := test.CreateUserAccountInternalServerError(t, context.Background(), service, ctrl, payloadInternalErr)

}