package main

import (
	"io/fs"
	"net/http"

	authApi "github.com/cryonayes/GoShare/api"
	fileApi "github.com/cryonayes/GoShare/api/file"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func Setup(app *fiber.App, fs fs.FS) {
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(fs),
	}))

	app.Post("/api/login", authApi.Login)
	app.Post("/api/register", authApi.Register)
	app.Post("/api/upload", fileApi.EndpointUploadFile)
}
