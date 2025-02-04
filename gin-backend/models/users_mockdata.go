package models

import (
	"fmt"
	"gorm.io/gorm"
)

func SeedMockUsers(db *gorm.DB) error {
	mockUsers := []User{
		{Name: "Alice Johnson", Email: "alice@example.com"},
		{Name: "Bob Smith", Email: "bob@example.com"},
		{Name: "Charlie Brown", Email: "charlie@example.com"},
		{Name: "David Miller", Email: "david@example.com"},
		{Name: "Eva Williams", Email: "eva@example.com"},
	}

	for _, user := range mockUsers {
		// Kullanıcı zaten var mı?
		var exists bool
		err := db.Model(&User{}).Where("email = ?", user.Email).Select("count(*) > 0").Scan(&exists).Error
		if err != nil {
			return fmt.Errorf("kullanıcı kontrolü hatası: %v", err)
		}

		// Yoksa ekle
		if !exists {
			err = db.Create(&user).Error
			if err != nil {
				return fmt.Errorf("mock user eklenemedi: %v", err)
			}
		}
	}
	fmt.Println("Users mock kullanıcılar başarıyla eklendi veya zaten mevcut!")
	return nil
}
