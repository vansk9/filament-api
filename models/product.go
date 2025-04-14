package models

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Returnable  bool      `json:"returnable"`
	Shippable   bool      `json:"shippable"`
	Inventory   Inventory `json:"inventory" gorm:"foreignKey:ProductID"`
}
