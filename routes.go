package main

import (
	authApi "github.com/cryonayes/GoShare/api"
	fileApi "github.com/cryonayes/GoShare/api/file"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {
	app.Use(cors.New(cors.ConfigDefault))
	app.Post("/api/login", authApi.Login)
	app.Post("/api/register", authApi.Register)
	app.Post("/api/upload", fileApi.EndpointUploadFile)
}
