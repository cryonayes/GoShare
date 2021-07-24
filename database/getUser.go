package database

import (
	models "github.com/cryonayes/GoShare/models"
	"github.com/cryonayes/GoShare/utils"
)

func GetUserFromUsername(username string) (*models.User, *utils.ApiError) {
	if connected := CheckConnection(); !connected {
		return nil, utils.NewJSONError(utils.DatabaseConnErr)
	}

	var user = &models.User{}
	err := DBConn.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, utils.NewJSONError(utils.UserNotFound)
	}

	return user, nil
}
