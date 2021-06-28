package file

import "github.com/gofiber/fiber/v2"

func EndpointUploadFile(c *fiber.Ctx) error {
	if _, err := c.MultipartForm(); err == nil {
		// TODO(Check if user is authenticated)
		// TODO(Validate and upload file with an unique name)
		// TODO(Add uploaded file's path to database with username associated with it)
		return c.SendString("Uploading is not available at the moment")
	}
	return c.SendString("Error while uploading")
}
