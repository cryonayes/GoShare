package routes

import (
	authApi "github.com/cryonayes/StajProje/api"
	fileApi "github.com/cryonayes/StajProje/api/file"
	userApi "github.com/cryonayes/StajProje/api/user"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Static("/", "./static")
	app.Post("/api/login", authApi.Login)
	app.Post("/api/register", userApi.EndpointRegister)
	app.Post("/api/upload", fileApi.EndpointUploadFile)
}
