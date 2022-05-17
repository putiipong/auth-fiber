package routes

import (
	"crud/types"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

const (
	jwtSecret = "eiei"
)

type (
	MsgLogin types.Login
	MsgToken types.Token
)

func Auth(c *fiber.Ctx) error {
	var l MsgLogin
	err := c.BodyParser(&l)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "please insert ",
		})
		return nil
	}
	if l.Username != "test" || l.Password != "password" {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
		return nil
	}
	token, err := createToken()
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "StatusInternalServerError",
			"msg":   err.Error(),
		})
		return nil
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"user":          l.Username,
	})
	return nil
}

func createToken() (MsgToken, error) {
	var msgToken MsgToken
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = utils.UUID()
	claims["name"] = "test test"
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return msgToken, err
	}
	msgToken.AccessToken = t

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = utils.UUID()
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	rt, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return msgToken, err
	}
	msgToken.RefreshToken = rt
	return msgToken, nil
}

//https://github.com/gofiber/jwt
func AuthorizationRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		// Filter:         nil,
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningKey:     []byte(jwtSecret),
		// SigningKeys:   nil,
		SigningMethod: "HS256",
		// ContextKey:    nil,
		// Claims:        nil,
		// TokenLookup:   nil,
		// AuthScheme:    nil,
	})
}

func AuthError(c *fiber.Ctx, e error) error {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
		"msg":   e.Error(),
	})
	return nil
}

func AuthSuccess(c *fiber.Ctx) error {
	c.Next()
	return nil
}

func Profile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"sub":  claims["sub"].(string),
		"name": claims["name"].(string),
	})
	return nil
}
