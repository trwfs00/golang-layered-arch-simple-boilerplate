package user

import (
	"github.com/gofiber/fiber/v2"
)

type CreateUserRequest struct {
	Name  string  `json:"name"`
	Phone *string `json:"phone"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	request := CreateUserRequest{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	err := h.createUserService.Execute(request.Name, request.Phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
