package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// .env dosyasından veritabanı bilgilerini al
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	// Veritabanına bağlan
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	// Bağlantıyı test et
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Bağlantı testi başarısız: %v", err)
	}
	fmt.Println("PostgreSQL'e başarıyla bağlanıldı!")

	// Tablo oluştur
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE
		);
	`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Tablo oluşturulamadı: %v", err)
	}
	fmt.Println("Tablo başarıyla oluşturuldu!")
}
