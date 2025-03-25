package routes

import "github.com/gofiber/fiber/v2"

func Index(c *fiber.Ctx) error {
    // Render index.html from ./public
    return c.Render("index", fiber.Map{})
}