package file

import (
	"encoding/json"
	"github.com/cryonayes/GoShare/api"
	"github.com/gofiber/fiber/v2"
)

func GetUploadedFiles(ctx *fiber.Ctx) error {
	loggedIn, email := api.CheckAuthentication(ctx)
	if !loggedIn || email == "" {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: "Please login!",
			Data:    nil,
		})
	}

	test := &struct{ Data string `json:"data"` }{"myData"}
	data, _ := json.Marshal(test)

	return ctx.JSON(api.Success{
		Success: true,
		Message: string(data),
		Data:    nil,
	})
}