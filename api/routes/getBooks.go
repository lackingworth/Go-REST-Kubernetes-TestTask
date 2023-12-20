package routes

import (
	"github.com/gofiber/fiber/v2"
	
	m "github.com/lackingworth/Go-REST-API/models"
)

func (r *Repository) GetBooks(c *fiber.Ctx) error {
	bookModels := &[]m.Books{}

	err := r.DB.Find(bookModels).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error":"Error while finding entries"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Books fetched successfully",
		"data": bookModels,
	})
}