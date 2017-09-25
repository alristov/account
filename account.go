package main

import (
	"fmt"
	"github.com/account/app"
	"github.com/account/db"

	"github.com/goadesign/goa"
)

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
	Repository db.Repository
}

// NewAccountController creates a account controller.
func NewAccountController(service *goa.Service, repo db.Repository) *AccountController {
	return &AccountController{Controller: service.NewController("AccountController"), Repository: repo}
}

// CreateUser runs the CreateUser action.
func (c *AccountController) CreateUser(ctx *app.CreateUserAccountContext) error {
	// AccountController_CreateUser: start_implement

	fmt.Println("Test1")

	// Put your logic here
	Account, err := c.Repository.CreateUser(ctx.Payload.Name, ctx.Payload.Username, ctx.Payload.Password, ctx.Payload.Email)
	// AccountController_CreateUser: end_implement

	if err != nil {
		ctx.InternalServerError(err)
	}
	res := &app.Account{
		Name:     Account.Name,
		Username: Account.Username,
		Email:    Account.Email,
		Password: Account.Password,
	}

	return ctx.Created(res)
}

func (c *AccountController) GetUser(ctx *app.GetUserAccountContext) error {

	Account, err := c.Repository.GetUser(ctx.AccountID)

	if err != nil {
		e := err.(*goa.ErrorResponse)

		switch e.Status {
		case 404:
			return ctx.NotFound()
		default:
			return ctx.InternalServerError(err)
		}
	}

	res := &app.Account{
		ID:       Account.ID,
		Name:     Account.Name,
		Username: Account.Username,
		Email:    Account.Email,
		Password: Account.Password,
	}

	return ctx.OK(res)
}

func (c *AccountController) UpdateUser(ctx *app.UpdateUserAccountContext) error {

	Account, err := c.Repository.UpdateUser(ctx.AccountID, ctx.Payload.Name, ctx.Payload.Username, ctx.Payload.Password, ctx.Payload.Email)

	if err != nil {
		ctx.InternalServerError(err)
	}

	if err != nil {
		e := err.(*goa.ErrorResponse)

		switch e.Status {
		case 404:
			return ctx.NotFound()
		default:
			return ctx.InternalServerError(err)
		}
	}

	res := &app.Account{
		ID:       Account.ID,
		Name:     Account.Name,
		Username: Account.Username,
		Email:    Account.Email,
		Password: Account.Password,
	}

	return ctx.OK(res)
}

func (c *AccountController) DeleteUser(ctx *app.DeleteUserAccountContext) error {
	err := c.Repository.DeleteUser(ctx.AccountID)

	if err != nil {
		e := err.(*goa.ErrorResponse)

		switch e.Status {
		case 204:
			return ctx.NoContent()
		case 404:
			return ctx.NotFound()
		default:
			return ctx.InternalServerError(err)
		}
	}

	return ctx.NoContent()
}
