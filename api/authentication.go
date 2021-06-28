package api

import (
	"github.com/cryonayes/StajProje/database"
	"github.com/cryonayes/StajProje/errors"
	"github.com/cryonayes/StajProje/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const SecretKey = "secret"

func Login(c *fiber.Ctx) error {
	var data models.User

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var userData = models.User{}
	database.DBConn.Where("username = ?", data.Username).First(&userData)

	if userData == (models.User{}) {
		err := errors.NewError(errors.UserNotFound)
		return c.JSON(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(errors.NewError(errors.InvalidCredentials))
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userData.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(errors.LoginFailed)
	}

	cookie := fiber.Cookie{
		Name:     "login",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "login-success",
	})

}

func Register(c *fiber.Ctx) error {
	var userRegister models.UserRegister

	if err := c.BodyParser(&userRegister); err != nil {
		mErr := errors.NewError(errors.RegisterFailed)
		return c.JSON(mErr)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(userRegister.Password), 14)

	user := models.User{}
	database.DBConn.Where("username = ?", userRegister.Username).First(&user)

	if user != (models.User{}) {
		return c.JSON(errors.NewError(errors.UserAlreadyExists))
	}

	user = models.User{
		Username: userRegister.Username,
		Password: string(password),
	}
	dbResponse := database.DBConn.Create(&user)

	if mErr := dbResponse.Error; mErr != nil {
		return c.JSON(errors.NewError(errors.RegisterFailed))
	}

	return c.JSON(user)
}
