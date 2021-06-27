package user

import (
	"github.com/gofiber/fiber/v2"
)

func EndpointRegister(c *fiber.Ctx) error {
	if form, err := c.MultipartForm(); err == nil {
		if usrName, passwd := form.Value[username], form.Value[password]; len(usrName) > 0 && len(passwd) > 0 {
			newUser := userData{
				Username: usrName[0],
				Password: passwd[0],
			}
			// TODO(Check if username already exists)
			// TODO(Insert data into database)
			return c.SendString(newUser.String())
		}
	}
	return c.SendString("Error while registering")
}
