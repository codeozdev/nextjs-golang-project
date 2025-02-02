package handlers

import (
	"awesomeProject/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetAllUsersHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := models.GetAllUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Veritabanı hatası"})
			return
		}

		if len(users) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users":   users,
			"status":  "200",
			"message": "Kullanıcılar başarıyla getirildi",
		})
	}
}

func GetUserByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
			return
		}

		user, err := models.GetUserByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Kullanıcı bulunamadı"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":    user,
			"status":  "200",
			"message": "Kullanıcı başarıyla getirildi",
		})
	}
}

func CreateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newUser.Name = strings.ToLower(newUser.Name)
		newUser.Email = strings.ToLower(newUser.Email)

		if !strings.Contains(newUser.Email, "@") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz email adresi"})
			return
		}

		err := models.CreateUser(db, &newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı eklenemedi"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "Kullanıcı başarıyla eklendi",
			"user":    newUser,
		})
	}
}

func UpdateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
			return
		}

		var updates map[string]interface{}
		if err := c.ShouldBindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = models.UpdateUser(db, id, updates)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Güncelleme hatası"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Kullanıcı başarıyla güncellendi",
		})
	}
}

func DeleteUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz ID"})
			return
		}

		err = models.DeleteUser(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Silme hatası"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Kullanıcı başarıyla silindi",
		})
	}
}
