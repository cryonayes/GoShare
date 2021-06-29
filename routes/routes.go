package routes

import (
	authApi "github.com/cryonayes/StajProje/api"
	fileApi "github.com/cryonayes/StajProje/api/file"
	"github.com/cryonayes/StajProje/views"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", views.ServeIndex)
	app.Get("/login", views.ServeLogin)
	app.Get("/register", views.ServeRegister)
	app.Get("/private", views.ServePrivate)

	app.Post("/api/login", authApi.Login)
	app.Post("/api/register", authApi.Register)
	app.Post("/api/upload", fileApi.EndpointUploadFile)
}
