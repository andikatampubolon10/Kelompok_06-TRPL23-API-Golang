package controller

import (
	"cbt-api/config"
	"cbt-api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetKursusById will return course details based on id_kursus
func GetKursusById(c *gin.Context) {
	idKursus := c.Param("id_kursus") // Ambil id_kursus dari parameter URL

	// Cari kursus berdasarkan id_kursus
	var kursus entity.Kursus
	if err := config.DB.Where("id_kursus = ?", idKursus).First(&kursus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching kursus"})
		return
	}

	// Jika kursus tidak ditemukan
	if (kursus == entity.Kursus{}) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Kursus not found for the given id_kursus"})
		return
	}

	// Kirimkan data kursus ke client
	c.JSON(http.StatusOK, gin.H{
		"message":  "Kursus retrieved successfully",
		"kursus":   kursus,
	})
}
