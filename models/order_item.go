package models

import "time"

type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Total     float64   `json:"total"` // <--- ini penting
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Product Product `json:"product" gorm:"constraint:OnDelete:CASCADE"`
}
