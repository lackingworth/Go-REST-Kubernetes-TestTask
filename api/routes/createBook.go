package routes

import (
	"github.com/gofiber/fiber/v2"
	
	m "github.com/lackingworth/Go-REST-API/models"
)

func (r *Repository) CreateBook(c *fiber.Ctx) error {
	book := m.Book{}

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"Error":"Error while parsing body"})
	}

	err := r.DB.Create(&book).Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error":"Error while creating new book entry"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Successfully added new book"})
}