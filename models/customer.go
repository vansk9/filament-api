package models

type Customer struct {
	ID uint `json:"id" gorm:"primaryKey"`
	// UserID  uint   `json:"user_id"` // foreign key
	Name    string `json:"name"`
	Email   string `json:"email"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
}
