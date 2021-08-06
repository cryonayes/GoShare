package file

import (
	"github.com/cryonayes/GoShare/api"
	"github.com/cryonayes/GoShare/database"
	appmodels "github.com/cryonayes/GoShare/models"
	"github.com/cryonayes/GoShare/utils"
	"github.com/gofiber/fiber/v2"
	"time"
)

func DownloadFile(ctx *fiber.Ctx) error {
	if dbconn := database.CheckConnection(); !dbconn {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.DatabaseConnErr,
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

	fileAccessToken := ctx.Params("accesstoken", "")
	if fileAccessToken != "" {
		var userFile = appmodels.FileModel{}
		database.DBConn.Table("file_models").
			Where("access_code = ? AND access_token = ?", accessCode, fileAccessToken).
			First(&userFile)

		if userFile.Shared && userFile.ShareTime.Unix() <= time.Now().Unix() {
			database.DBConn.Model(userFile).Where("access_token = ?", fileAccessToken).Updates(map[string]interface{}{
				"shared":       false,
				"share_time":   time.Now(),
				"access_token": "",
			})

			return ctx.JSON(api.Failure{
				Success: false,
				Message: utils.FileShareExpired,
				Data:    nil,
			})
		}

		if userFile.Shared && userFile.ShareTime.Unix() > time.Now().Unix() {
			return ctx.SendFile("./uploads/"+userFile.HashedFileName, true)
		}
	}

	loggedIn, email := api.CheckAuthentication(ctx)
	if !loggedIn || email == "" {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.PleaseLogin,
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

	return ctx.SendFile("./uploads/"+fileModel.HashedFileName, true)
}
