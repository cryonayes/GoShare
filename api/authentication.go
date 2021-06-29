package api

import (
	"github.com/cryonayes/StajProje/database"
	"github.com/cryonayes/StajProje/errorUtil"
	"github.com/cryonayes/StajProje/models"
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
	var data models.User

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var userData = models.User{}
	database.DBConn.Where("username = ?", data.Username).First(&userData)

	if userData == (models.User{}) {
		err := errorUtil.NewError(errorUtil.UserNotFound)
		return c.JSON(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(errorUtil.NewError(errorUtil.InvalidCredentials))
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Username: userData.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(errorUtil.LoginFailed)
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
	var userRegister models.UserRegister

	if err := c.BodyParser(&userRegister); err != nil {
		mErr := errorUtil.NewError(errorUtil.RegisterFailed)
		return c.JSON(mErr)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(userRegister.Password), 14)

	user := models.User{}
	database.DBConn.Where("username = ?", userRegister.Username).First(&user)

	if user != (models.User{}) {
		return c.JSON(errorUtil.NewError(errorUtil.UserAlreadyExists))
	}

	user = models.User{
		Username: userRegister.Username,
		Password: string(password),
	}
	dbResponse := database.DBConn.Create(&user)

	if mErr := dbResponse.Error; mErr != nil {
		return c.JSON(errorUtil.NewError(errorUtil.RegisterFailed))
	}

	return c.JSON(user)
}

func CheckAuthentication(ctx *fiber.Ctx) bool {
	token := ctx.Cookies("token", "")
	if token == "" {
		return false
	}
	claims := &Claims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		ctx.ClearCookie("token")
		return false
	}

	if !jwtToken.Valid {
		ctx.ClearCookie("token")
		return false
	}

	return true
}
