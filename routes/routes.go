package routes

import (
	file2 "github.com/cryonayes/StajProje/api/file"
	user2 "github.com/cryonayes/StajProje/api/user"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Static("/", "./static")
	app.Post("/api/register", user2.EndpointRegister)
	app.Post("/api/upload", file2.EndpointUploadFile)

}