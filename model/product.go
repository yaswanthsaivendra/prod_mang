package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName        string  `gorm:"size:255; not null" json:"product_name"`
	ProductDescription string  `gorm:"type:text; not null" json:"product_description"`
	ProductPrice       float64 `gorm:"not null" json:"product_price"`
	UserID             uint
	Images             []Image
}

type Image struct {
	ProductImage           string `gorm:"not null" json:"product_image"`
	CompressedProductImage string `json:"compressed_product_image"`
}
