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

func DeleteCustomer(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Only admin can delete customers"})
	}
	customerID := c.Params("id")
	if customerID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Customer ID is required"})
	}

	var customer models.Customer
	if err := config.DB.First(&customer, customerID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}
	if err := config.DB.Delete(&customer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete customer"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Customer deleted successfully"})
}
