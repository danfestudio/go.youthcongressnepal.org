package routes

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	
    // Render login.html from ./public
    return c.Render("login", fiber.Map{})
    
}