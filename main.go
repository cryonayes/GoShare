package main

import (
	"fmt"
	"github.com/cryonayes/StajProje/routes"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()
	routes.Setup(app)

	err := app.Listen(":8080")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Cannot initialize server!")
		return
	}
}
