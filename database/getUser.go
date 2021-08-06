package database

import (
	models "github.com/cryonayes/GoShare/models"
	"github.com/cryonayes/GoShare/utils"
)

func GetUserFromEmail(email string) (*models.User, error) {
	if connected := CheckConnection(); !connected {
		return nil, utils.NewError(utils.DatabaseConnErr)
	}

	var user = &models.User{}
	err := DBConn.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, utils.NewError(utils.UserNotFound)
	}

	return user, nil
}
