package main

import (
	"io/fs"

	authApi "github.com/cryonayes/GoShare/api"
	fileApi "github.com/cryonayes/GoShare/api/file"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App, fs fs.FS) {

	/* app.Use("/", filesystem.New(filesystem.Config{
		Root:  http.FS(fs),
		Index: "/register.html",
	})) */

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://localhost:3000, http://localhost:3000, http://localhost:8080, https://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/api/login", authApi.Login)

	app.Post("/api/register", authApi.Register)
	app.Post("/api/upload", fileApi.EndpointUploadFile)
}
