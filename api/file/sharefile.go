package file

import (
	"github.com/cryonayes/GoShare/api"
	"github.com/cryonayes/GoShare/database"
	appmodels "github.com/cryonayes/GoShare/models"
	"github.com/cryonayes/GoShare/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func ShareFile(ctx *fiber.Ctx) error {
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

	type fileShareDatas struct {
		AccessCode string `json:"accesscode"`
		ShareTime string `json:"sharetime"`
	}

	var fileData = fileShareDatas{}
	err := ctx.BodyParser(&fileData)
	if err != nil{
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.InvalidFileCode,
			Data:    nil,
		})
	}

	shareTime := fileData.ShareTime
	if shareTime == "" {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.InvalidTimeValue,
			Data:    nil,
		})
	}

	var userFile appmodels.FileModel
	database.DBConn.Table("file_models").Where("access_code = ?", fileData.AccessCode).Find(&userFile)

	if userFile.Owner != email {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.Unauthorized,
			Data:    nil,
		})
	}

	shareTimeInt, err := strconv.ParseInt(shareTime, 10, 64)
	if err != nil {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.InvalidTimeValue,
			Data:    nil,
		})
	}

	convertedTime := time.Unix(shareTimeInt, 0)
	if convertedTime.Unix() < time.Now().Unix() {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.InvalidTimeValue,
			Data:    nil,
		})
	}

	userFile.Shared = true
	userFile.ShareTime = convertedTime
	userFile.AccessToken = utils.GetMD5String(userFile.ShareTime.String() + userFile.AccessCode)

	database.DBConn.Table("file_models").Where("access_code = ?", fileData.AccessCode).Updates(&userFile)

	return ctx.JSON(api.Success{
		Success: true,
		Message: utils.FileShared,
		Data: struct {
			AccessLink string `json:"accesslink"`
		}{
			userFile.AccessCode + "/" +userFile.AccessToken,
		},
	})
}
