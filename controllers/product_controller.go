package controllers

import (
	"filament-api/config"
	"filament-api/models"
	"filament-api/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	var input struct {
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Returnable  bool    `json:"returnable"`
		Shippable   bool    `json:"shippable"`

		Inventory struct {
			SKU           string `json:"sku" validate:"required"`
			Barcode       string `json:"barcode"`
			Stock         int    `json:"stock" validate:"required"`
			SecurityStock int    `json:"security_stock"`
		} `json:"inventory" validate:"required"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := utils.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Returnable:  input.Returnable,
		Shippable:   input.Shippable,
	}

	inventory := models.Inventory{
		SKU:           input.Inventory.SKU,
		Barcode:       input.Inventory.Barcode,
		Stock:         input.Inventory.Stock,
		SecurityStock: input.Inventory.SecurityStock,
	}

	// Simpan product + inventory via relasi
	product.Inventory = inventory

	if err := config.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}
