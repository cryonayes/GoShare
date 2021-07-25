package main

import (
	"embed"
	"fmt"
	"github.com/cryonayes/GoShare/database"
	"io/fs"
	"log"
	"os"

	"github.com/cryonayes/GoShare/api/file"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//go:embed nextjs/dist
//go:embed nextjs/dist/_next
//go:embed nextjs/dist/_next/static/chunks/pages/*.js
//go:embed nextjs/dist/_next/static/*/*.js
var nextFS embed.FS

func main() {

	distFS, err := fs.Sub(nextFS, "nextjs/dist")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(file.UploadDir); os.IsNotExist(err) {
		err := os.Mkdir(file.UploadDir, os.ModeType)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Cannot create uploads directory!")
		}
	}

	app := fiber.New(fiber.Config{})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://localhost:3000, http://localhost:3000, http://localhost:8080, https://localhost:8080",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	Setup(app, distFS)

	database.Connect()

	err = app.Listen(":8080")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Cannot initialize server!")
		return
	}
}
