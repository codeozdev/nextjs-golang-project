package main

import (
	"awesomeProject/database"
	"awesomeProject/handlers"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// Veritabanı bağlantısını başlat
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Veritabanı bağlantısı başarısız: %v", err)
	}

	// Mock verileri ekle
	err = models.SeedMockUsers(db)
	if err != nil {
		log.Fatalf("Users mock veriler eklenemedi: %v", err)
	}

	err = models.SeedMockProducts(db)
	if err != nil {
		log.Fatalf("Products mock veriler eklenemedi: %v", err)
	}

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.SetTrustedProxies([]string{"192.168.1.0/24", "127.0.0.1"})

	// Route'ları tanımla
	r.GET("/users", handlers.GetAllUsersHandler(db))
	r.GET("/users/:id", handlers.GetUserByIDHandler(db))
	r.POST("/users", handlers.CreateUserHandler(db))
	r.PATCH("/users/:id", handlers.UpdateUserHandler(db))
	r.DELETE("/users/:id", handlers.DeleteUserHandler(db))

	// Route'ları tanımla
	r.GET("/products", handlers.GetAllProductsHandler(db))
	//r.GET("/products/:id", handlers.GetUserByIDHandler(db))
	//r.POST("/products", handlers.CreateUserHandler(db))
	//r.PATCH("/products/:id", handlers.UpdateUserHandler(db))
	//r.DELETE("/products/:id", handlers.DeleteUserHandler(db))

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
