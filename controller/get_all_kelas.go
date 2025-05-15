package controller

import (
	"cbt-api/config"
	"cbt-api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllKelas - Menampilkan seluruh kelas
func GetAllKelas(c *gin.Context) {
	// Inisialisasi slice untuk menampung semua kelas
	var kelasList []entity.Kelas

	// Query untuk mengambil seluruh data kelas
	err := config.DB.Find(&kelasList).Error

	// Jika terjadi error
	if err != nil {
		// Jika ada kesalahan dalam mengambil data
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan saat mengambil data kelas"})
		return
	}

	// Kirimkan response seluruh kelas dalam format JSON
	c.JSON(http.StatusOK, gin.H{
		"message": "Kelas retrieved successfully",
		"kelas":   kelasList,
	})
}
