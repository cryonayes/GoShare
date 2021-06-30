package api

import (
	"github.com/cryonayes/StajProje/database"
	"github.com/cryonayes/StajProje/models"
	"github.com/cryonayes/StajProje/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var SecretKey = []byte("secret")

type Claims struct {
	Username string
	jwt.StandardClaims
}

func Login(c *fiber.Ctx) error {
	if connected := database.CheckConnection(); !connected {
		return c.JSON(utils.NewJSONError(utils.DatabaseConnErr))
	}
	var data models.User

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(utils.NewJSONError(utils.RequestError))
	}

	if data.Username == "" || data.Password == "" {
		return c.JSON(utils.NewJSONError(utils.RequestError))
	}

	userData, err := database.GetUserFromUsername(data.Username)
	if err != nil {
		return c.JSON(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(utils.NewJSONError(utils.InvalidCredentials))
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: userData.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	token, tErr := claims.SignedString(SecretKey)
	if tErr != nil {
		c.Status(fiber.StatusInternalServerError)
		c.ClearCookie("token")
		return c.JSON(utils.InternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "token",
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
	if connected := database.CheckConnection(); !connected {
		return utils.NewError(utils.DatabaseConnErr)
	}
	var userRegister models.UserRegister

	if err := c.BodyParser(&userRegister); err != nil {
		mErr := utils.NewJSONError(utils.RegisterFailed)
		return c.JSON(mErr)
	}
	if userRegister.Username == "" {
		return c.JSON(Failure{
			Success: false,
			Message: utils.InvalidUsername,
			Data:    nil,
		})
	}
	if userRegister.Password == "" {
		return c.JSON(Failure{
			Success: false,
			Message: utils.InvalidPassword,
			Data:    nil,
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(userRegister.Password), 14)

	user := models.User{}
	database.DBConn.Where("username = ?", userRegister.Username).First(&user)

	if user != (models.User{}) {
		return c.JSON(utils.NewJSONError(utils.UserAlreadyExists))
	}

	user = models.User{
		Username: userRegister.Username,
		Password: string(password),
	}
	dbResponse := database.DBConn.Create(&user)

	if mErr := dbResponse.Error; mErr != nil {
		return c.JSON(utils.NewJSONError(utils.RegisterFailed))
	}

	return c.JSON(user)
}

func CheckAuthentication(ctx *fiber.Ctx) (bool, string) {
	token := ctx.Cookies("token", "")
	if token == "" {
		return false, ""
	}
	claims := &Claims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil || !jwtToken.Valid {
		ctx.ClearCookie("token")
		return false, ""
	}
	return true, claims.Username
}
