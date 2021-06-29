package views

import (
	"github.com/gofiber/fiber/v2"
)

func ServeIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"testData": "TEST",
	})
}
