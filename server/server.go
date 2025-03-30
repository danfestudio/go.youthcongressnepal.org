package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func StartServer() {	
	
	engine := html.New("./public", ".html")
	app := fiber.New(fiber.Config{
		Views: engine, 
	})

	Routes(app)

	app.Static("/static", "./static")
	app.Static("/assets", "./assets")	
	
	app.Listen(":8001")
}