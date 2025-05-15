package controller

import (
	"cbt-api/config"
	"cbt-api/entity"
	"golang.org/x/crypto/bcrypt"  // Used for hashing and checking password
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginUjian handles login to the exam (Ujian) using the password provided by the user
func AccessUjian(c *gin.Context) {
	// Get the id_ujian and password from the request parameters
	idUjian := c.Param("id_ujian") // Retrieve exam ID from URL
	var userInput struct {
		PasswordMasuk string `json:"password_masuk"` // The password input by the user
	}
	
	// Bind the incoming JSON to the userInput struct
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Retrieve the Ujian (exam) from the database based on id_ujian
	var ujian entity.Ujian
	if err := config.DB.Where("id_ujian = ?", idUjian).First(&ujian).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ujian not found"})
		return
	}

	// Compare the hashed password stored in the database with the input password
	err := bcrypt.CompareHashAndPassword([]byte(ujian.PasswordMasuk), []byte(userInput.PasswordMasuk))
	if err != nil {
		// If the password does not match
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return
	}

	// If the password is correct, send a success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"ujian":   ujian.NamaUjian,
		"id_ujian": ujian.IdUjian,
	})
}
