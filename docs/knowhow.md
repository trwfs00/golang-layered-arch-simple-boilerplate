# How to Create a New API Endpoint

This guide outlines the steps to add a new API endpoint to the existing Go Fiber application. The process includes creating a handler, service, and repository, while following the layered architecture.

## Step 1: Define the API Specification

Before implementing the new API endpoint, define what the endpoint will do, including:

- **Endpoint URL**: The path where the API will be accessed.
- **HTTP Method**: The HTTP method (GET, POST, PUT, DELETE, etc.).
- **Request Body**: The structure of the request data (if applicable).
- **Response Structure**: The expected response data.

### Example Specification

For example, let's create an API endpoint to update a user's information:

- **Endpoint URL**: /api/v1/user/:id
- **HTTP Method**: PUT
- **Request Body**:

```json
{
  "username": "newusername",
  "email": "newemail@example.com"
}
```

- **Response Structure**:

```json
{
  "message": "User updated successfully"
}
```

### Step 2: Create the Handler

In the `api/handler/user/` directory, create a new file named `update.go`. This file will contain the HTTP handler for the new endpoint.

#### Example Handler Code

`api/handler/user/update.go`

```go
package user

import (
	"github.com/gofiber/fiber/v2"
	"boilerplate/api/service/user/command"
)

type UpdateUserRequest struct {
	Name    string  `json:"name"`
	Phone   *string `json:"phone"`
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	request := UpdateUserRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	err := h.updateUserService.Execute(id, request.Name, request.Phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}
```

### Step 2.1: Update handler dependencies

In the `api/handler/user/` directory, update a file named `user.go`. This file will contain all user handler dependencies.

`api/handler/user/user.go`

```go
package user

import (
	"boilerplate/api/service/user/command"
	"boilerplate/api/service/user/query"
)

type UserHandler struct {
	getUserByIdService query.GetUserByIdService
	createUserService  command.CreateUserService
    updateUserService command.UpdateUserService // add this line of code to append new dependency type
}

// declare an updateUserService to parameters
func NewUserHandler(getUserByIdService query.GetUserByIdService, createUserService command.CreateUserService, updateUserService command.UpdateUserService) *UserHandler {
	return &UserHandler{
		getUserByIdService: getUserByIdService,
		createUserService:  createUserService,
        updateUserService: updateUserService, // add this line of code to append new dependency
	}
}
```

### Step 3: Create the Service

In the `api/service/user/command/` directory, create a new file named `update.go`. This file will implement the business logic for updating a user.

#### Example Service Code

`api/service/user/command/update.go`

```go
package command

import (
	"boilerplate/api/repository"
)

type UpdateUserService interface {
	Execute(id string, name string, phone *string) error
}

type updateUserService struct {
	userRepo repository.UserRepository
}

func NewUpdateUserService(repo repository.UserRepository) UpdateUserService {
	return &updateUserService{userRepo: repo}
}

func (s *updateUserService) Execute(id string, name string, phone *string) error {
	// Assuming ID is of type string; you might need to convert it to the appropriate type
	user := &repository.User{
		Name:     name,
		Phone:    phone,
	}
	return s.userRepo.UpdateUser(id, user)
}
```

### Step 4: Update the Repository

In the `api/repository/users.go` file, add the method for updating a user.

#### Example Repository Code

`api/repository/users.go`

```go
package repository

import (
	"fmt"
	"gorm.io/gorm"
    "boilerplate/lib/database/entity"
)

type userRepository struct {
	db *gorm.DB
}

// Other methods...

func (r *userRepository) UpdateUser(id string, user *entity.User) error {
	if err := r.db.Model(&entity.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}
```

### Step 5: Update the Router

In the `api/router.go` file, add the new route for the `UpdateUser` endpoint.
```go
package api

import (
	"github.com/gofiber/fiber/v2"
	"boilerplate/api/handler/user"
)

func SetupRoutes(app *fiber.App, userHandler *user.UserHandler) {
	api := app.Group("/api/v1")

	api.Group("/user")
	{
		api.Get("/:id", userHandler.GetUserById)
		api.Post("/", userHandler.CreateUser)
		api.Put("/:id", userHandler.UpdateUser) // Add this line for update user
	}
}
```

### Step 6: Update Dependency Injection
In the `main.go` file, ensure the new service and handler are instantiated and passed to the router.

#### Example Dependency Injection Code

`main.go`

```go
package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"user/api"
	"user/api/handler/user"
	"user/api/repository"
	"user/api/service/user/command"
	"user/lib/database"
	"user/lib/environment"
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
    updateUserService := command.NewUpdateUserService(userRepo) // New service for updating user

	// Handlers
	userHandler := user.NewUserHandler(
        getUserByIdService,
        createUserService,
        updateUserService, // Pass update service to handler
    )

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
```

### Step 7: Test the New Endpoint
You can test the new update user endpoint with a curl command:
```curl
curl -X PUT http://localhost:3000/api/v1/user/1 \
-H "Content-Type: application/json" \
-d '{"name": "updatedusername", "phone": "001-002-0003"}'
```

### Summary
By following these steps, you can create a new API endpoint while adhering to the existing architecture. The process involves defining the endpoint, implementing the handler, service, and repository, updating the router, and ensuring proper dependency injection.

### Conclusion
This guide can be used as a reference whenever you need to add new API functionality to your Go Fiber project. Keep the API specifications clear, and always ensure to test the endpoints after implementing them.