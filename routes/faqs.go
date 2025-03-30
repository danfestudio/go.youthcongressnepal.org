package routes

import "github.com/gofiber/fiber/v2"

func FAQs(c *fiber.Ctx) error {
    
    // Render faqs.html from ./public
    return c.Render("faqs", fiber.Map{})
    
}