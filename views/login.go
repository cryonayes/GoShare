package views

import "github.com/gofiber/fiber/v2"

func ServeLogin(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{})
}
