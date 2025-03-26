package server

import (
	"github.com/danfelab/youthcongressnepal/routes"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	app.Get("/", routes.Index)

	app.Get("/register", routes.Register)
	app.Post("/register", routes.RegisterForm)

	app.Get("/login", routes.Login)
	
}
