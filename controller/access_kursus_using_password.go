package controller

import (
	"cbt-api/config"
	"cbt-api/entity"
	"golang.org/x/crypto/bcrypt" // Untuk hashing dan membandingkan password
	"github.com/gin-gonic/gin"
	"net/http"
)

// AccessKursus handles access to a course using the password provided by the user
func EnrollKursus(c *gin.Context) {
	// Get the id_kursus and password from the request parameters
	idKursus := c.Param("id_kursus") // Ambil id_kursus dari URL
	var userInput struct {
		Password string `json:"password"` // Password yang diinput oleh pengguna
	}

	// Bind the incoming JSON to the userInput struct
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Retrieve the Kursus (course) from the database based on id_kursus
	var kursus entity.Kursus
	if err := config.DB.Where("id_kursus = ?", idKursus).First(&kursus).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kursus not found"})
		return
	}

	// Compare the hashed password stored in the database with the input password
	err := bcrypt.CompareHashAndPassword([]byte(kursus.Password), []byte(userInput.Password))
	if err != nil {
		// If the password does not match
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	// If the password is correct, send a success response
	c.JSON(http.StatusOK, gin.H{
		"message":   "Access granted to the course",
		"kursus":    kursus.NamaKursus,
		"id_kursus": kursus.IdKursus,
	})
}
