package database

import (
	"awesomeProject/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// .env dosyasından veritabanı bilgilerini al
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Veritabanına bağlan
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("veritabanına bağlanılamadı: %v", err)
	}

	fmt.Println("PostgreSQL'e başarıyla bağlanıldı!")

	// Otomatik Migration
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return nil, fmt.Errorf("tablo oluşturulamadı: %v", err)
	}
	fmt.Println("Tablo başarıyla oluşturuldu!")

	return DB, nil
}
