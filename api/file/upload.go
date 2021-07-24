package file

import (
	"fmt"
	"github.com/cryonayes/GoShare/api"
	"github.com/cryonayes/GoShare/database"
	models "github.com/cryonayes/GoShare/models"
	"github.com/cryonayes/GoShare/utils"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

const (
	UploadDir = "./uploads"
)

func EndpointUploadFile(ctx *fiber.Ctx) error {
	if connected := database.CheckConnection(); !connected {
		return ctx.JSON(utils.NewJSONError(utils.DatabaseConnErr))
	}

	authenticated, username := api.CheckAuthentication(ctx)
	if !authenticated {
		return ctx.JSON(api.Failure{Success: false, Message: utils.Unauthenticated, Data: nil})
	}

	file, err := ctx.FormFile("document")
	if err != nil {
		return ctx.JSON(api.Failure{Success: false, Message: utils.UploadError, Data: nil})
	}

	fType, validErr := utils.CheckFileType(file)
	if validErr != nil || fType == "" {
		return ctx.JSON(utils.NewJSONError(utils.InvalidFileType))
	}

	uploadedTime := time.Now()
	hashedName := utils.GetMD5String(file.Filename + uploadedTime.String())

	uploadedFile := models.FileModel{
		OrigFileName:   file.Filename,
		HashedFileName: hashedName + "." + fType,
		FileType:       fType,
		FileSize:       file.Size,
		Owner:          username,
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
		Message: "File uploaded",
		Data:    uploadedFile,
	})
}
