package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

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

	// Tüm kullanıcıları getir
	r.GET("/users", func(c *gin.Context) {
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

	r.Run(":8080")

}
