package main

import (
	"awesomeProject/database"
	"awesomeProject/handlers"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
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
	database.InitDB()

	// Mock verileri ekle
	err := models.SeedMockUsers(database.DB)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.SetTrustedProxies([]string{"192.168.1.0/24", "127.0.0.1"})

	// Route'ları tanımla
	r.GET("/users", handlers.GetAllUsersHandler(database.DB))
	r.GET("/users/:id", handlers.GetUserByIDHandler(database.DB))
	r.POST("/users", handlers.CreateUserHandler(database.DB))
	r.PATCH("/users/:id", handlers.UpdateUserHandler(database.DB))
	r.DELETE("/users/:id", handlers.DeleteUserHandler(database.DB))

	r.Run(":8080")
}
