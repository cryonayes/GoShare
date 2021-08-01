package main

import (
	"fmt"
	"github.com/cryonayes/GoShare/database"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {

	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024, // 50MB limit for uploads
	})

	Setup(app)

	database.Connect()

	err := app.Listen(":21942")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Cannot initialize server!")
		return
	}
}
