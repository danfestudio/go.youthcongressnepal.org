package server

import (
	"github.com/danfelab/youthcongressnepal/routes"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	app.Get("/", routes.Index)
	app.Get("/assets/", routes.Assets)
}
