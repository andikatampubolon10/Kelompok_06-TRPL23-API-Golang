package controller

import (
	"cbt-api/config"
	"cbt-api/entity"
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetUjianById - Get Ujian details by idUjian
func GetUjianById(c *gin.Context) {
	// Get the idUjian from the URL parameter
	idUjian := c.Param("idUjian")

	// Initialize the Ujian variable to hold the result
	var ujian entity.Ujian

	// Query the database to find the Ujian with the given idUjian
	err := config.DB.Where("id_ujian = ?", idUjian).First(&ujian).Error

	// If there's an error while fetching the data
	if err != nil {
		// If no Ujian found or any other error occurs
		c.JSON(http.StatusNotFound, gin.H{"error": "Ujian not found"})
		return
	}

	// Send back the Ujian data as a JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "Ujian retrieved successfully",
		"ujian":   ujian,
	})
}
