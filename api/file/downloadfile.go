package file

import (
	"github.com/cryonayes/GoShare/api"
	"github.com/cryonayes/GoShare/database"
	appmodels "github.com/cryonayes/GoShare/models"
	"github.com/cryonayes/GoShare/utils"
	"github.com/gofiber/fiber/v2"
)

func DownloadFile(ctx *fiber.Ctx) error {
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

	accessCode := ctx.Params("accesscode", "")
	if accessCode == "" {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.InvalidFileCode,
			Data:    nil,
		})
	}

	var fileModel = appmodels.FileModel{}
	database.DBConn.Table("file_models").Where("access_code = ?", accessCode).First(&fileModel)
	if fileModel.HashedFileName == "" {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.FileNotExists,
			Data:    nil,
		})
	}

	if fileModel.Owner != email {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.Unauthorized,
			Data:    nil,
		})
	}

	return ctx.SendFile("./uploads/" + fileModel.HashedFileName)
}
