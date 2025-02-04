package handlers

import (
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllProductsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := models.GetAllProducts(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "Veritabanı hatası"})
			return
		}

		if len(products) == 0 {
			c.JSON(404, gin.H{"error": "Ürün bulunamadı"})
			return
		}

		c.JSON(200, gin.H{
			"products": products,
			"status":   "200",
			"message":  "Ürünler başarıyla getirildi",
		})

	}

}
