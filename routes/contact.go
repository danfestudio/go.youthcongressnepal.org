package routes

import "github.com/gofiber/fiber/v2"

func Contact(c *fiber.Ctx) error {
    
    // Render index.html from ./public
    return c.Render("contact", fiber.Map{})
    
}