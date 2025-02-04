package models

import (
	"fmt"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

func SeedMockProducts(db *gorm.DB) error {
	mockProducts := []Product{
		{Name: "iPhone 13 Pro", Price: 1399.99, Stock: 10},
		{Name: "Galaxy S22 Ultra", Price: 1199.49, Stock: 20},
		{Name: "Pixel 7 Pro", Price: 999.99, Stock: 30},
		{Name: "OnePlus 11", Price: 849.00, Stock: 40},
		{Name: "Xiaomi 13 Pro", Price: 749.50, Stock: 40},
	}

	for _, product := range mockProducts {
		var exists bool
		err := db.Model(&Product{}).Where("name = ?", product.Name).Select("count(*) > 0").Scan(&exists).Error
		if err != nil {
			return fmt.Errorf("bu ürün zaten var: %v", err)
		}

		if !exists {
			err = db.Create(&product).Error
			if err != nil {
				return fmt.Errorf("mock ürün eklenemedi: %v", err)
			}
		}
	}

	fmt.Println("Products mock ürünler başarıyla eklendi veya zaten mevcut!")
	return nil
}
