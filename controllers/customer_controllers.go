package controllers

import (
	"filament-api/config"
	"filament-api/models"
	"filament-api/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateCustomer(c *fiber.Ctx) error {
	customer := new(models.Customer)
	if err := c.BodyParser(customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := utils.Validate.Struct(customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	config.DB.Create(&customer)
	return c.Status(fiber.StatusCreated).JSON(customer)
}
