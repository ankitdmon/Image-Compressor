package models

import "time"

type Product struct {
	ProductID          int       `json:"product_id" gorm:"primaryKey"`
	ProductName        string    `json:"product_name"`
	ProductDescription string    `json:"product_description"`
	ProductImages      []string  `json:"product_images" gorm:"type:type:VARCHAR(255)[]"`
	ProductPrice       float64   `json:"product_price"`
	CompressedImages   []string  `json:"compressed_product_images" gorm:"type:type:VARCHAR(255)[]"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
