package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	type User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var users = []User{
		{ID: 1, Name: "Ahmet", Email: "ahmet@example.com"},
		{ID: 2, Name: "Mehmet", Email: "mehmet@example.com"},
		{ID: 3, Name: "Ayşe", Email: "ayse@example.com"},
	}

	r := gin.Default()
	r.Use(CORSMiddleware()) // CORS Middleware’i ekle

	// ************ GET ************
	// Tüm kullanıcıları getir
	r.GET("/users", func(c *gin.Context) {

		if len(users) == 0 {
			c.JSON(404, gin.H{
				"error": "Kullanıcı bulunamadı",
			})
			return
		}

		c.JSON(200, gin.H{
			"users":   users,
			"status":  "success",
			"message": "Kullanıcılar başarıyla getirildi",
		})
	})

	// ID'ye göre kullanıcı getir
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			if fmt.Sprint(user.ID) == id {
				c.JSON(200, gin.H{
					"user": user,
				})
				return
			}
		}
		c.JSON(404, gin.H{
			"error": "Kullanıcı bulunamadı",
		})
	})

	// ************ POST ************
	r.POST("/users", func(c *gin.Context) {
		var newUser User

		// Gelen JSON verisini newUser struct'ına bağla
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Yeni kullanıcıya bir ID ata ve listeye ekle
		newUser.ID = len(users) + 1
		users = append(users, newUser)
		c.JSON(201, gin.H{
			"status":  "success",
			"message": "Kullanıcı başarıyla eklendi",
			//"user":    newUser, // Eklenen kullanıcıyı döndür (isteğe bağlı)
		})
	})

	r.Run(":8080")

}
