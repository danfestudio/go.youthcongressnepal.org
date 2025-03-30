package routes

import "github.com/gofiber/fiber/v2"

func Congratulation(c *fiber.Ctx) error {
    
    // Render index.html from ./public
    return c.Render("congratulation", fiber.Map{})
    
}