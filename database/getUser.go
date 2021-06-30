package database

import (
	"github.com/cryonayes/StajProje/errorUtil"
	"github.com/cryonayes/StajProje/models"
)

func GetUserFromUsername(username string) (*models.User, *errorUtil.ApiError) {
	if connected := CheckConnection(); !connected {
		return nil, errorUtil.NewJSONError(errorUtil.DatabaseConnErr)
	}

	var user = &models.User{}
	err := DBConn.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, errorUtil.NewJSONError(errorUtil.UserNotFound)
	}

	return user, nil
}
