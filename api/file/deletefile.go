package file

import (
	"github.com/cryonayes/GoShare/api"
	"github.com/cryonayes/GoShare/database"
	appmodels "github.com/cryonayes/GoShare/models"
	"github.com/cryonayes/GoShare/utils"
	"github.com/gofiber/fiber/v2"
	"os"
)

func DeleteFile(ctx *fiber.Ctx) error {
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

	var fileShareData = appmodels.FileShareDatas{}
	err := ctx.BodyParser(&fileShareData)
	if err != nil {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.InvalidFileCode,
			Data:    nil,
		})
	}

	var userFile appmodels.FileModel
	database.DBConn.Table("file_models").Where("access_code = ?", fileShareData.AccessCode).First(&userFile)
	if userFile.Owner != email {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.Unauthorized,
			Data:    nil,
		})
	}

	err = os.Remove("./uploads/" + userFile.HashedFileName)
	if err != nil {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.ErrorWhileDeleting,
			Data:    nil,
		})
	}

	dbExec := database.DBConn.Table("file_models").Where("access_code = ?", fileShareData.AccessCode).Delete(userFile)
	if dbExec.Error != nil {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.ErrorWhileDeleting,
			Data:    nil,
		})
	}

	return ctx.JSON(api.Success{
		Success: true,
		Message: utils.FileDeleted,
		Data:    nil,
	})
}
