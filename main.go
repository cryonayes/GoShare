package main

import (
	"fmt"
	"github.com/cryonayes/StajProje/api/file"
	"github.com/cryonayes/StajProje/database"
	"github.com/cryonayes/StajProje/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"os"
)

func main() {
	database.Connect()

	if _, err := os.Stat(file.UploadDir); os.IsNotExist(err) {
		err := os.Mkdir(file.UploadDir, os.ModeType)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Cannot create uploads directory!")
		}
	}

	app := fiber.New(fiber.Config{
		// Initialize html template engine
		Views: html.New("./public", ".html"),
	})

	routes.Setup(app)

	err := app.Listen(":8080")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Cannot initialize server!")
		return
	}
}
