package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()
	app.Static("/", "./static")
	err := app.Listen(":8080")
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Cannot initialize server!")
		return
	}
}
