package user

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	user, err := h.getUserByIdService.Execute(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}
