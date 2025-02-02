package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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
	r.Use(CORSMiddleware())                                      // CORS Middleware’i ekle
	r.SetTrustedProxies([]string{"192.168.1.0/24", "127.0.0.1"}) // Güvenilir proxy'leri ayarla

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
			"status":  "200",
			"message": "Kullanıcılar başarıyla getirildi",
		})
	})

	// ID'ye göre kullanıcı getir (parametreden veriyi aldik)
	r.GET("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam) // string'i integer'a çevir
		if err != nil {
			c.JSON(400, gin.H{
				"error":   "Geçersiz ID",
				"message": "ID bir sayı olmalıdır",
			})
			return
		}

		for _, user := range users {
			if user.ID == id {
				c.JSON(200, gin.H{
					"user":    user,
					"status":  "200",
					"message": "Kullanıcı başarıyla getirildi",
				})
				return
			}
		}

		// Kullanıcı bulunamazsa
		c.JSON(404, gin.H{
			"error":   "Kullanıcı bulunamadı",
			"message": fmt.Sprintf("%d ID'li kullanıcı bulunamadı", id),
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

	// ************ PUT ************
	r.PATCH("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updates map[string]interface{}

		if err := c.ShouldBindJSON(&updates); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		for i, user := range users {
			if fmt.Sprint(user.ID) == id {
				// Sadece belirtilen alanları güncelle
				for key, value := range updates {
					switch key {
					case "name":
						users[i].Name = value.(string)
					case "email":
						users[i].Email = value.(string)
					}
				}
				c.JSON(200, gin.H{
					"status":  "success",
					"message": "Kullanıcı başarıyla güncellendi",
					"user":    users[i],
				})
				return
			}
		}

		c.JSON(404, gin.H{"error": "Kullanıcı bulunamadı"})
	})

	r.Run(":8080")

}
