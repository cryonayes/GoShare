package views

import "github.com/gofiber/fiber/v2"

func ServeRegister(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{})
}
