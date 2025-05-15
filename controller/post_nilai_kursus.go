package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"cbt-api/config"
	"cbt-api/entity"
	"strconv"
)

func PostNilaiKursus(c *gin.Context) {
	// Extract id_kursus and id_siswa from URL parameters
	idKursus := c.Param("id_kursus")
	idSiswa := c.Param("id_siswa")

	// Define the structure of the request body
	var input struct {
		NilaiTipeUjian float64 `json:"nilai_tipe_ujian"`
		IdTipeUjian     uint64  `json:"id_tipe_ujian"`
	}

	// Bind the request body to the input structure
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "detail": err.Error()})
		return
	}

	// Convert id_kursus and id_siswa to uint64 (assuming the IDs are numeric)
	kursusID, err := strconv.ParseUint(idKursus, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_kursus", "detail": err.Error()})
		return
	}

	siswaID, err := strconv.ParseUint(idSiswa, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_siswa", "detail": err.Error()})
		return
	}

	// Create a new NilaiKursus instance with the provided data
	nilaiKursus := entity.NilaiKursus{
		NilaiTipeUjian: input.NilaiTipeUjian,
		IdKursus:       kursusID,
		IdSiswa:        siswaID,
		IdTipeUjian:    input.IdTipeUjian,
	}

	// Insert the new NilaiKursus into the database
	if err := config.DB.Create(&nilaiKursus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the score", "detail": err.Error()})
		return
	}

	// Return a success response with the newly created NilaiKursus data
	c.JSON(http.StatusOK, gin.H{
		"message":      "Nilai kursus added successfully",
		"nilai_kursus": nilaiKursus,
	})
}
