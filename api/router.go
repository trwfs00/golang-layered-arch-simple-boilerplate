package api

import (
	"boilerplate/api/handler/user"

	"github.com/gofiber/fiber/v2"
)

// Set up routes
func SetupRoutes(app *fiber.App, userHandler *user.UserHandler) {
	api := app.Group("/api/v1")

	userGroup := api.Group("/user")
	userGroup.Get("/:id", userHandler.GetUserById)
	userGroup.Post("/", userHandler.CreateUser)

	// TODO: Add more routes
}
