package model

import (
	"github.com/yaswanthsaivendra/prod_mang/database"
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
	gorm.Model
	ProductImage           string `gorm:"not null" json:"product_image"`
	CompressedProductImage string `json:"compressed_product_image"`
	ProductID              uint
}

func (product *Product) Save() (*Product, error) {
	err := database.Database.Create(&product).Error
	if err != nil {
		return &Product{}, err
	}
	return product, nil
}
