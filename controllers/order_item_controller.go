package controllers

import (
	"filament-api/config"
	"filament-api/models"
	"filament-api/utils"

	"github.com/gofiber/fiber/v2"
)

type AddProductToOrderInput struct {
	OrderID   uint `json:"order_id" validate:"required"`
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,min=1"`
}

func AddProductToOrder(c *fiber.Ctx) error {
	var input AddProductToOrderInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := utils.Validate.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	// Ambil produk dan harga
	var product models.Product
	if err := config.DB.First(&product, input.ProductID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	price := product.Price
	quantity := input.Quantity
	total := price * float64(quantity)

	orderItem := models.OrderItem{
		OrderID:   input.OrderID,
		ProductID: input.ProductID,
		Quantity:  quantity,
		Price:     price,
		Total:     total,
		Product:   product,
	}

	if err := config.DB.Create(&orderItem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add product to order"})
	}

	return c.Status(fiber.StatusCreated).JSON(orderItem)
}

func GetOrderItems(c *fiber.Ctx) error {
	var items []models.OrderItem

	if err := config.DB.
		Preload("Product").
		Preload("Product.Inventory").
		Find(&items).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch order items",
		})
	}

	return c.JSON(items)
}

func GetOrderItemByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.OrderItem

	if err := config.DB.
		Preload("Product").
		Preload("Product.Inventory").
		First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order item not found",
		})
	}

	return c.JSON(item)
}

func DeleteOrderItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.OrderItem

	if err := config.DB.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order item not found",
		})
	}

	if err := config.DB.Delete(&item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete order item",
		})
	}

	return c.JSON(fiber.Map{"message": "Order item deleted successfully"})
}
