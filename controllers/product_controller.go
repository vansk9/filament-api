package controllers

import (
	"filament-api/config"
	"filament-api/models"
	"filament-api/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Only admin can create products"})
	}

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

	product.Inventory = inventory

	if err := config.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Only admin can delete products"})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
	}

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	if err := config.DB.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product deleted successfully"})
}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if err := config.DB.Preload("Inventory").First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	if err := config.DB.Preload("Inventory").Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}
	return c.JSON(products)
}
