package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func StartServer() {
	
	// Initialize Fiber with the HTML template engine
	engine := html.New("./public", ".html")
	app := fiber.New(fiber.Config{
		Views: engine, // Set the template engine
	})

	Routes(app)

	app.Static("/static", "./static")
	app.Static("/assets", "./assets")	
	
	// Start the server
	app.Listen(":3000")
}