package user

import (
	"github.com/cryonayes/StajProje/models"
	"github.com/gofiber/fiber/v2"
)

func EndpointRegister(c *fiber.Ctx) error {
	if form, err := c.MultipartForm(); err == nil {
		if usrName, passwd := form.Value[models.Username], form.Value[models.Password]; len(usrName) > 0 && len(passwd) > 0 {
			newUser := models.User{
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
