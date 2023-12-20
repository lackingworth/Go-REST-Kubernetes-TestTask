package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	m "github.com/lackingworth/Go-REST-API/models"
)

func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	bookModel := m.Books{}
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error":"Invalid ID, ID cannot be empty"})
	}

	err := r.DB.Delete(bookModel, id)

	if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error":"Error while deleting entry"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Book deleted successfully"})
}