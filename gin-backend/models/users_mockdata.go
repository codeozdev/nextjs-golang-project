package models

import (
	"database/sql"
	"fmt"
)

func SeedMockUsers(db *sql.DB) error {
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
		err := db.QueryRow(`SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`, user.Email).Scan(&exists)
		if err != nil {
			return fmt.Errorf("kullanıcı kontrolü hatası: %v", err)
		}

		// Yoksa ekle
		if !exists {
			err = CreateUser(db, &user)
			if err != nil {
				return fmt.Errorf("mock user eklenemedi: %v", err)
			}
		}
	}
	fmt.Println("Mock kullanıcılar başarıyla eklendi veya zaten mevcut!")
	return nil
}
