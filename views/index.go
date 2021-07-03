package views

import (
	authApi "github.com/cryonayes/StajProje/api"
	"github.com/gofiber/fiber/v2"
)

func ServeIndex(c *fiber.Ctx) error {

	exists, username := authApi.CheckAuthentication(c)
	if !exists && username == "" {
		return c.Render("index", fiber.Map{})
	}

	// TODO(Get user's files from database)

	return c.Render("userpanel", fiber.Map{
		"data": map[string][]string{"test": {"a", "b"}},
	})
}
