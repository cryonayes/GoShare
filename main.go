package main

import (
	"fmt"
	"github.com/cryonayes/GoShare/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func main() {

	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://localhost:3000, http://localhost:3000, http://localhost:8080, https://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	Setup(app)

	database.Connect()

	err := app.Listen(":21942")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Cannot initialize server!")
		return
	}
}
