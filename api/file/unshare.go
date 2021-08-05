package file

import (
	"github.com/cryonayes/GoShare/api"
	"github.com/cryonayes/GoShare/database"
	appmodels "github.com/cryonayes/GoShare/models"
	"github.com/cryonayes/GoShare/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

func UnshareFile(ctx *fiber.Ctx) error {
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

	var fileData = appmodels.FileAccessCode{}
	err := ctx.BodyParser(&fileData)
	if err != nil {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.InvalidFileCode,
			Data:    nil,
		})
	}

	var userFile appmodels.FileModel
	database.DBConn.Table("file_models").Where("access_code = ?", fileData.AccessCode).First(&userFile)
	if userFile.Owner != email {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.Unauthorized,
			Data:    nil,
		})
	}

	dbExec := database.DBConn.Table("file_models").Where("access_code = ?", fileData.AccessCode).Updates(map[string]interface{}{
		"shared":     false,
		"share_time": time.Now(),
	})

	if dbExec.Error != nil {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.ErrorWhileUnshare,
			Data:    nil,
		})
	}

	return ctx.JSON(api.Success{
		Success: true,
		Message: utils.FileUpdated,
		Data:    nil,
	})
}
