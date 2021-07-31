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
	// TODO(Convert to api failure)
	if connected := database.CheckConnection(); !connected {
		return ctx.JSON(utils.NewJSONError(utils.DatabaseConnErr))
	}

	authenticated, userEmail := api.CheckAuthentication(ctx)
	if !authenticated {
		return ctx.JSON(api.Failure{Success: false, Message: utils.Unauthenticated, Data: nil})
	}

	file, err := ctx.FormFile("testFile")
	if err != nil {
		return ctx.JSON(api.Failure{Success: false, Message: utils.UploadError, Data: nil})
	}

	// TODO(Convert to api failure)
	fType, validErr := utils.CheckFileType(file)
	if validErr != nil || fType == "" {
		return ctx.JSON(utils.NewJSONError(utils.InvalidFileType))
	}

	uploadedTime := time.Now()
	hashedFileName := utils.GetMD5String(file.Filename + uploadedTime.String())
	// hashedUserMail := utils.GetMD5String(userEmail)

	uploadedFile := models.FileModel{
		OrigFileName:   file.Filename,
		HashedFileName: hashedFileName + "." + fType,
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
		Data: struct {
			OrigFileName   string    `json:"orig_file_name"`
			FileType       string    `json:"file_type"`
			FileSize       int64     `json:"file_size"`
			Owner          string    `json:"owner"`
			CreationDate   time.Time `json:"creation_date"`
		}{
			uploadedFile.OrigFileName,
			uploadedFile.FileType,
			uploadedFile.FileSize,
			uploadedFile.Owner,
			uploadedFile.CreationDate,
		},
	})
}
