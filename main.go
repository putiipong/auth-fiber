package main

import (
	"crud/routes"

	fiber "github.com/gofiber/fiber/v2"
)

var (
	app = fiber.New()
)

func main() {

	//not AuthorizationRequired
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/login", routes.Auth)

	//AuthorizationRequired Action
	app.Use(routes.AuthorizationRequired())

	//need AuthorizationRequired
	app.Get("/profile", routes.Profile)
	//end AuthorizationRequired
	app.Listen(":3000")

}
