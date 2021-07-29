package file

import (
	"github.com/cryonayes/GoShare/api"
	"github.com/cryonayes/GoShare/database"
	app_models "github.com/cryonayes/GoShare/models"
	"github.com/gofiber/fiber/v2"
)

func GetUploadedFiles(ctx *fiber.Ctx) error {
	loggedIn, email := api.CheckAuthentication(ctx)
	if !loggedIn || email == "" {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: "Please login!",
			Data:    nil,
		})
	}

	var user = app_models.User{}
	var userFiles []app_models.UserFileModel

	database.DBConn.Table("users").Where("email = ?", email).First(&user)
	database.DBConn.Table("file_models").Where("owner = ?", email).Find(&userFiles)

	var userData = app_models.UserDataModel{
		Name: user.Name,
		Lastname: user.LastName,
		Files: userFiles,
	}

	return ctx.JSON(api.Success{
		Success: true,
		Message: "Success",
		Data:    userData,
	})
}