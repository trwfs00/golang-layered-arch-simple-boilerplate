package main

import (
	"fmt"
	"log"

	"boilerplate/api"
	"boilerplate/api/handler/user"
	"boilerplate/api/repository"
	"boilerplate/api/service/user/command"
	"boilerplate/api/service/user/query"
	"boilerplate/lib/database"
	"boilerplate/lib/environment"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	environment.New(0) // Pass 0 if the env file is in the current directory

	// Get DSN from environment variables
	dsn := environment.GetString(environment.DsnKey)
	if dsn == "" {
		log.Fatal("DB_DSN not set in environment")
	}

	// Initialize Database
	database.InitDatabaseWithDSN(dsn)

	// Dependency Injection
	// Repos
	userRepo := repository.NewUserRepository(database.DB)

	// Services
	// user query
	getUserByIdService := query.NewGetUserByIdService(userRepo)

	// user command
	createUserService := command.NewCreateUserService(userRepo)

	// Handlers
	userHandler := user.NewUserHandler(getUserByIdService, createUserService)

	// Set up Fiber
	app := fiber.New()
	api.SetupRoutes(app, userHandler)

	// Get the port from the environment
	port := environment.GetString(environment.ServicePort)
	if port == "" {
		port = "9090" // Default port if not set
	}

	// Start the server
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
