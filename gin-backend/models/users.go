package models

import (
	"gorm.io/gorm"
	"strings"
)

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
}

//	*************** CRUD OPERATIONS ***************
//

// GetAllUsers tüm kullanıcıları getirir
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	return users, err
}

// GetUserByID ID'ye göre kullanıcı getirir
func GetUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser yeni kullanıcı ekler
func CreateUser(db *gorm.DB, user *User) error {
	user.Name = strings.ToLower(user.Name)
	user.Email = strings.ToLower(user.Email)
	return db.Create(user).Error
}

// UpdateUser kullanıcıyı günceller
func UpdateUser(db *gorm.DB, id uint, updates map[string]interface{}) error {
	return db.Model(&User{}).Where("id = ?", id).Updates(updates).Error
}

// DeleteUser kullanıcıyı siler
func DeleteUser(db *gorm.DB, id uint) error {
	return db.Delete(&User{}, id).Error
}
