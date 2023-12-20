package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	m "github.com/lackingworth/Go-REST-API/models"
)

func (r *Repository) GetBookByID(c *fiber.Ctx) error {
	bookModel := &m.Books{}
	id := c.Params("id")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error":"Invalid ID, ID cannot be empty"})
	}

	err := r.DB.Where("id = ?", id).First(bookModel).Error

	if err != nil  && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error":"Error while finding entry"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Book found successfully",
		"data": bookModel,
	})
}