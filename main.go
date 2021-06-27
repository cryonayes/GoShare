package main

import (
	"fmt"
	"github.com/cryonayes/StajProje/file"
	"github.com/cryonayes/StajProje/user"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()
	app.Static("/", "./static")
	app.Post("/api/register", user.EndpointRegister)
	app.Post("/api/upload", file.EndpointUploadFile)

	err := app.Listen(":8080")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Cannot initialize server!")
		return
	}
}
