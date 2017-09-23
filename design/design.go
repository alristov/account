package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("account", func() {
	Title("Goa user")
	Description("Goa user description")
	Scheme("http")
	Host("localhost:8080")

	/*
		ResponseTemplate(Created, func(pattern string) {
			Description("Resource created")
			Status(201)
			Headers(func() {
				Header("Location", String, "href to created resource", func() {
					Pattern(pattern)
				})
			})
		})
	*/

})

var _ = Resource("account", func() {

	DefaultMedia(Account)
	BasePath("/accounts")

	Action("CreateUser", func() {
		Routing(
			POST(""),
		)
		Description("Create new account")
		Payload(func() {
			Member("name")
			Member("username")
			Member("email")
			Member("password")
			Required("name", "username", "email", "password")
		})

		Response(Created, Account)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

	Action("GetUser", func() {
		Routing(
			GET("/:accountID"),
		)
		Description("Get account by ID")
		Params(func() {
			Param("accountID", Integer, "Account ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("DeleteUser", func() {
		Routing(
			DELETE("/:accountID"),
		)
		Params(func() {
			Param("accountID", Integer, "Account ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("UpdateUser", func() {
		Routing(
			PUT("/:accountID"),
		)
		Description("change account attributes")
		Params(func() {
			Param("accountID", Integer, "Account ID")
		})
		Payload(func() {
			Member("name")
			Member("username")
			Member("email")
			Member("password")
			Required("name")
			Required("username")
			Required("email")
			Required("password")
		})
		Response(NotFound)
		Response(OK, Account)
		Response(BadRequest, ErrorMedia)
		Response(InternalServerError, ErrorMedia)
	})

})

var Account = MediaType("application/vnd.account+json", func() {
	Description("Media type of account!!!")

	Attributes(func() {

		Attribute("id", Integer, "id of account", func() {
			Example(1)
		})
		Attribute("name", String, "Name of account", func() {
			Example("testName")
		})
		Attribute("username", String, "Username of account", func() {
			Example("testUsername")
		})
		Attribute("email", String, "Email of account", func() {
			Format("email")
			Example("me@goa.design:)")
		})
		Attribute("password", String, "Password of account", func() {
			Example("testPassword")
		})

		Required("id", "name", "username", "email", "password")
	})

	View("default", func() {
		Attribute("id", Integer, "Account ID")
		Attribute("name", String, "Account Name")
		Attribute("username", String, "Account Username")
		Attribute("email", String, "Account Email")
		Attribute("password", String, "Account Password")
	})

})
