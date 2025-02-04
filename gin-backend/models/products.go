package models

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	Stock     int       `json:"stock" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// *************** CRUD OPERATIONS ***************

// GetAllProducts tüm ürünleri getirir
func GetAllProducts(db *gorm.DB) ([]Product, error) {
	var products []Product
	err := db.Find(&products).Error
	return products, err
}
