package main

import (
	authApi "github.com/cryonayes/GoShare/api"
	fileApi "github.com/cryonayes/GoShare/api/file"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(app *fiber.App) {

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://localhost:3000, http://localhost:3000, http://localhost:8080, https://localhost:8080",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Content-Length, Accept, Cookie, X-TOKEN",
	}))

	app.Post("/api/login", authApi.Login)
	app.Post("/api/register", authApi.Register)
	app.Post("/api/authcheck", authApi.AuthCheckForFrontend)

	app.Post("/api/upload", fileApi.EndpointUploadFile)
	app.Post("/api/files", fileApi.GetUploadedFiles)
	app.Post("/api/share", fileApi.ShareFile)
	app.Post("/api/delete", fileApi.DeleteFile)
	app.Post("/api/unshare", fileApi.UnshareFile)
	app.Get("/api/download/:accesscode/:accesstoken?", fileApi.DownloadFile)
}
