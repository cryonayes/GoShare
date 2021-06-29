package views

import (
	"github.com/cryonayes/StajProje/api"
	"github.com/gofiber/fiber/v2"
)

func ServePrivate(ctx *fiber.Ctx) error {
	authenticated := api.CheckAuthentication(ctx)
	if !authenticated {
		return ctx.SendString("You are not authenticated!")
	}
	return ctx.Render("private", &fiber.Map{})
}
