package routes

import "github.com/gofiber/fiber/v2"

func About(c *fiber.Ctx) error {
    
    // Render index.html from ./public
    return c.Render("about", fiber.Map{})
    
}