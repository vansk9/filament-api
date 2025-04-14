package main

import (
	"filament-api/config"
	"filament-api/controllers"
	"filament-api/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	app := fiber.New()

	api := app.Group("/api")
	api.Post("/login", controllers.Login)
	api.Post("/register", controllers.Register)

	auth := api.Group("/auth", middleware.JWTProtected())
	auth.Post("/customers", controllers.CreateCustomer)          // Membuat customer baru
	auth.Post("/products", controllers.CreateProduct)            // Membuat produk baru
	auth.Post("/orders", controllers.CreateOrder)                // Membuat order baru (mendapatkan order id untuk digunakan di order item)
	auth.Post("/order-items", controllers.AddProductToOrder)     // Menambahkan Produk ke dalam order
	auth.Delete("/order-items/:id", controllers.DeleteOrderItem) // Menghapus produk dari order

	log.Fatal(app.Listen(":3000"))
}
