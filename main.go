package main

import (
	"crud/configs"
	"crud/db"
	"crud/middleware"
	"crud/routes"

	fiber "github.com/gofiber/fiber/v2"
)

var (
	app = fiber.New()
)

func main() {

	// //not AuthorizationRequired
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	// app.Post("/login", routes.Auth)

	// //AuthorizationRequired Action
	// app.Use(routes.AuthorizationRequired())

	// //need AuthorizationRequired
	// app.Get("/profile", routes.Profile)
	// //end AuthorizationRequired
	// app.Listen(":3000")

	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.

	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).

	db.StartServer(app)
}
