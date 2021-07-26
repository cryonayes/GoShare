package main

import (
	"fmt"
	"github.com/cryonayes/GoShare/database"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {

	app := fiber.New(fiber.Config{})

	Setup(app)

	database.Connect()

	err := app.Listen(":21942")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Cannot initialize server!")
		return
	}
}
