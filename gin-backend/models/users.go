package models

import (
	"database/sql"
	"fmt"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser Yeni kullanıcı ekle
func CreateUser(db *sql.DB, user *User) error {
	user.Name = strings.ToLower(user.Name)
	user.Email = strings.ToLower(user.Email)

	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	return db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

// GetAllUsers Tüm kullanıcıları getir
func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByID ID'ye göre kullanıcı getir
func GetUserByID(db *sql.DB, id int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser Kullanıcı güncelle
func UpdateUser(db *sql.DB, id int, updates map[string]interface{}) error {
	for key, value := range updates {
		query := fmt.Sprintf("UPDATE users SET %s = $1 WHERE id = $2", key)
		_, err := db.Exec(query, value, id)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteUser Kullanıcı sil
func DeleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
