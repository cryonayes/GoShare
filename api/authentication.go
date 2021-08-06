package api

import (
	"time"

	"github.com/cryonayes/GoShare/database"
	appmodels "github.com/cryonayes/GoShare/models"
	"github.com/cryonayes/GoShare/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte("secret")

type Claims struct {
	Email string
	jwt.StandardClaims
}

func Login(c *fiber.Ctx) error {
	if connected := database.CheckConnection(); !connected {
		return c.JSON(Failure{Success: false, Message: utils.DatabaseConnErr, Data: nil})
	}
	var data appmodels.User

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(Failure{Success: false, Message: utils.RequestError, Data: nil})
	}

	if data.Email == "" || data.Password == "" {
		return c.JSON(Failure{Success: false, Message: utils.RequestError, Data: nil})
	}

	userData, err := database.GetUserFromEmail(data.Email)
	if err != nil {
		return c.JSON(Failure{Success: false, Message: utils.UserNotFound, Data: nil})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password)); err != nil {
		return c.JSON(Failure{Success: false, Message: utils.InvalidCredentials, Data: nil})
	}

	expirationDate := time.Now().Add(time.Hour * 8)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Email: userData.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationDate.Unix(),
		},
	})

	token, tErr := claims.SignedString(SecretKey)
	if tErr != nil {
		c.Status(fiber.StatusInternalServerError)
		c.ClearCookie("token")
		return c.JSON(Failure{Success: false, Message: utils.InternalServerError, Data: nil})
	}

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expirationDate,
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(Success{
		Success: true,
		Message: "Login Success",
		Data:    nil,
	})
}

func Register(ctx *fiber.Ctx) error {
	if dbconn := database.CheckConnection(); !dbconn {
		return ctx.JSON(Failure{
			Success: false,
			Message: utils.DatabaseConnErr,
			Data:    nil,
		})
	}

	var userRegister appmodels.UserRegister
	if err := ctx.BodyParser(&userRegister); err != nil {
		return ctx.JSON(Failure{Success: false, Message: utils.RegisterFailed, Data: nil})
	}

	if userRegister.Email == "" {
		return ctx.JSON(Failure{
			Success: false,
			Message: utils.InvalidEmail,
			Data:    nil,
		})
	}
	if userRegister.Password == "" {
		return ctx.JSON(Failure{
			Success: false,
			Message: utils.InvalidPassword,
			Data:    nil,
		})
	}
	if userRegister.Name == "" {
		return ctx.JSON(Failure{
			Success: false,
			Message: utils.InvalidName,
			Data:    nil,
		})
	}
	if userRegister.LastName == "" {
		return ctx.JSON(Failure{
			Success: false,
			Message: utils.InvalidLastname,
			Data:    nil,
		})
	}
	if userRegister.Password != userRegister.PasswordRepeat {
		return ctx.JSON(Failure{
			Success: false,
			Message: utils.PasswordRepeatWrong,
			Data:    nil,
		})
	}

	user := appmodels.User{}
	database.DBConn.Where("email = ?", userRegister.Email).First(&user)

	if user != (appmodels.User{}) {
		return ctx.JSON(Failure{Success: false, Message: utils.UserAlreadyExists, Data: nil})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(userRegister.Password), 14)
	user = appmodels.User{
		Email:    userRegister.Email,
		Password: string(password),
		Name:     userRegister.Name,
		LastName: userRegister.LastName,
	}

	dbResponse := database.DBConn.Create(&user)

	if mErr := dbResponse.Error; mErr != nil {
		return ctx.JSON(Failure{Success: false, Message: utils.RegisterFailed, Data: nil})
	}

	return ctx.JSON(Success{
		true,
		utils.UserRegistered,
		nil,
	})
}

func CheckAuthentication(ctx *fiber.Ctx) (bool, string) {

	var token = ctx.Cookies("token", "")

	if token == "" {
		token = string(ctx.Request().Header.Peek("X-TOKEN"))
	}

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

	if claims.ExpiresAt <= time.Now().Unix() {
		return false, ""
	}

	return true, claims.Email
}

func AuthCheckForFrontend(ctx *fiber.Ctx) error {
	auth, mail := CheckAuthentication(ctx)

	if auth || mail != "" {
		return ctx.JSON(Success{
			Success: true,
			Message: utils.Authenticated,
			Data:    nil,
		})
	} else {
		return ctx.JSON(Failure{
			Success: false,
			Message: utils.Unauthenticated,
			Data:    nil,
		})
	}
}
