package controllers

import (
	"filament-api/config"
	"filament-api/models"
	"filament-api/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type OrderRequest struct {
	CustomerName  string `json:"customer_name" validate:"required"`
	Status        string `json:"status" validate:"required,oneof=New Processing Shipped Delivered Cancelled"`
	Currency      string `json:"currency" validate:"required"`
	Country       string `json:"country" validate:"required"`
	StreetAddress string `json:"street_address" validate:"required"`
	City          string `json:"city" validate:"required"`
	State         string `json:"state" validate:"required"`
	Zip           string `json:"zip" validate:"required"`
}

func CreateOrder(c *fiber.Ctx) error {
	var input struct {
		CustomerName  string `json:"customer_name" validate:"required"`
		Status        string `json:"status" validate:"required,oneof=New Processing Shipped Delivered Cancelled"`
		Currency      string `json:"currency" validate:"required"`
		Country       string `json:"country" validate:"required"`
		StreetAddress string `json:"street_address" validate:"required"`
		City          string `json:"city" validate:"required"`
		State         string `json:"state" validate:"required"`
		Zip           string `json:"zip" validate:"required"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validasi input tidak kosong
	if err := utils.Validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"validation_error": err.Error()})
	}

	// Cari customer berdasarkan name
	var customer models.Customer
	if err := config.DB.Where("name = ?", input.CustomerName).First(&customer).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Customer dengan nama tersebut tidak ditemukan",
		})
	}

	// Generate order number
	orderNumber := fmt.Sprintf("ORD-%d", time.Now().UnixNano())

	order := models.Order{
		OrderNumber:   orderNumber,
		CustomerID:    customer.ID,
		Status:        input.Status,
		Currency:      input.Currency,
		Country:       input.Country,
		StreetAddress: input.StreetAddress,
		City:          input.City,
		State:         input.State,
		Zip:           input.Zip,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Preload customer agar response lengkap
	if err := config.DB.Preload("Customer").First(&order, order.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
