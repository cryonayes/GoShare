package file

import (
	"fmt"
	"github.com/cryonayes/GoShare/api"
	"github.com/cryonayes/GoShare/database"
	appmodels "github.com/cryonayes/GoShare/models"
	"github.com/cryonayes/GoShare/utils"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
	"time"
)

const (
	UploadDir = "./uploads"
)

func EndpointUploadFile(ctx *fiber.Ctx) error {
	if dbconn := database.CheckConnection(); !dbconn {
		return ctx.JSON(api.Failure{
			Success: false,
			Message: utils.DatabaseConnErr,
			Data:    nil,
		})
	}

	authenticated, userEmail := api.CheckAuthentication(ctx)
	if !authenticated {
		return ctx.JSON(api.Failure{Success: false, Message: utils.Unauthenticated, Data: nil})
	}

	file, err := ctx.FormFile("fileupload")
	if err != nil {
		return ctx.JSON(api.Failure{Success: false, Message: utils.UploadError, Data: nil})
	}

	fType, validErr := utils.CheckFileType(file)
	if validErr != nil || fType == "" {
		return ctx.JSON(api.Failure{Success: false, Message: utils.InvalidFileType, Data: nil})
	}

	uploadedTime := time.Now()
	hashedFileName := utils.GetMD5String(file.Filename + uploadedTime.String())
	accessCode := utils.GetMD5String(hashedFileName) + utils.GetMD5String(strconv.FormatInt(file.Size*uploadedTime.Unix(), 10))

	uploadedFile := appmodels.FileModel{
		OrigFileName:   file.Filename,
		HashedFileName: hashedFileName + "." + fType,
		AccessCode:     accessCode,
		FileType:       fType,
		FileSize:       file.Size,
		Owner:          userEmail,
		IsEncrypted:    false,
		CreationDate:   time.Now(),
	}
	// TODO(Create unique access code for external access)
	err = ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", uploadedFile.HashedFileName))
	if err != nil {
		return ctx.JSON(api.Failure{Success: false, Message: utils.FileSavingError, Data: nil})
	}

	dbResponse := database.DBConn.Create(&uploadedFile)
	if mErr := dbResponse.Error; mErr != nil {
		err := os.Remove(fmt.Sprintf(UploadDir + "/" + uploadedFile.HashedFileName))
		if err != nil {
			return ctx.JSON(utils.NewJSONError(utils.UploadError))
		}
		return ctx.JSON(utils.NewJSONError(utils.DatabaseConnErr))
	}

	return ctx.JSON(&api.Success{
		Success: true,
		Message: "File uploaded!",
		Data: appmodels.UserFileModel{
			OrigFileName: uploadedFile.OrigFileName,
			FileType:     uploadedFile.FileType,
			FileSize:     uploadedFile.FileSize,
			Owner:        uploadedFile.Owner,
			CreationDate: uploadedFile.CreationDate,
		},
	})
}
