package main

import (
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"io/fs"
	"net/http"

	authApi "github.com/cryonayes/GoShare/api"
	fileApi "github.com/cryonayes/GoShare/api/file"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App, fs fs.FS) {

	app.Use("/", filesystem.New(filesystem.Config{
		Root:  http.FS(fs),
		Index: "/index.html",
		NotFoundFile: "/404.html",
	}))

	app.Use(cors.New(cors.ConfigDefault))
	app.Post("/api/login", authApi.Login)
	app.Post("/api/register", authApi.Register)
	app.Post("/api/upload", fileApi.EndpointUploadFile)
}
