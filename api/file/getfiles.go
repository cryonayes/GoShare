package file

import (
	"github.com/cryonayes/GoShare/api"
	"github.com/cryonayes/GoShare/database"
	appmodels "github.com/cryonayes/GoShare/models"
	"github.com/cryonayes/GoShare/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUploadedFiles(ctx *fiber.Ctx) error {
	if dbconn := database.CheckConnection(); !dbconn {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.DatabaseConnErr,
			Data:    nil,
		})
	}

	loggedIn, email := api.CheckAuthentication(ctx)
	if !loggedIn || email == "" {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.PleaseLogin,
			Data:    nil,
		})
	}

	var user = appmodels.User{}
	var userFiles []appmodels.UserFileModel

	database.DBConn.Table("users").Where("email = ?", email).First(&user)
	database.DBConn.Table("file_models").Where("owner = ?", email).Find(&userFiles)

	var userData = appmodels.UserDataModel{
		Name:     user.Name,
		Lastname: user.LastName,
		Files:    userFiles,
	}

	return ctx.JSON(api.Success{
		Success: true,
		Message: "Success",
		Data:    userData,
	})
}
