package file

import (
	"fmt"
	"github.com/cryonayes/StajProje/api"
	"github.com/cryonayes/StajProje/errorUtil"
	"github.com/cryonayes/StajProje/models"
	"github.com/gofiber/fiber/v2"
)

const (
	UploadDir = "./uploads"
)

func EndpointUploadFile(ctx *fiber.Ctx) error {
	authenticated := api.CheckAuthentication(ctx)
	if !authenticated {
		return ctx.JSON(api.Failure{Success: false, Message: errorUtil.Unauthenticated, Data: nil})
	}

	file, err := ctx.FormFile("document")
	if err != nil {
		return ctx.JSON(api.Failure{Success: false, Message: errorUtil.UploadError, Data: nil})
	}

	// TODO(Hash filename)
	err = ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	if err != nil {
		return ctx.JSON(api.Failure{Success: false, Message: errorUtil.FileSavingError, Data: nil})
	}

	return ctx.JSON(&api.Success{
		Success: true,
		Message: "File uploaded",
		Data: models.FileModel{
			FileName: file.Filename,
			FileSize: file.Size,
		},
	})
}
