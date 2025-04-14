package models

type Inventory struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	ProductID     uint   `json:"product_id"`
	SKU           string `json:"sku"`
	Barcode       string `json:"barcode"`
	Stock         int    `json:"stock"`
	SecurityStock int    `json:"security_stock"`

	// Gunakan pointer & json:"-" untuk menghindari infinite loop saat serialisasi
	Product *Product `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}
