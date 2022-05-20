package db

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Build Fiber connection URL.

	// Run server.
	if err := a.Listen(":8000"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
