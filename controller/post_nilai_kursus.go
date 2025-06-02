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
		IdTipeUjian uint64 `json:"id_tipe_ujian"`
	}

	// Bind the request body to the input structure
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "detail": err.Error()})
		return
	}

	// Convert id_kursus and id_siswa to uint64
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

	// Step 1: Sum all nilai from tipe_nilai table with the same id_tipe_ujian and id_siswa
	var totalNilai float64
	if err := config.DB.Model(&entity.TipeNilai{}).
		Where("id_tipe_ujian = ? AND id_siswa = ?", input.IdTipeUjian, siswaID).
		Select("COALESCE(SUM(nilai), 0)").
		Scan(&totalNilai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total nilai", "detail": err.Error()})
		return
	}

	// Step 2: Get the percentage from persentase table based on id_tipe_ujian and id_kursus
	var persentase entity.Persentase
	if err := config.DB.Where("id_tipe_ujian = ? AND id_kursus = ?", input.IdTipeUjian, kursusID).
		First(&persentase).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Persentase not found for the given id_tipe_ujian and id_kursus", "detail": err.Error()})
		return
	}

	// Step 3: Calculate the final score (total nilai * percentage / 100)
	finalScore := totalNilai * (persentase.Persentase / 100)

	// Step 4: Check if record already exists in nilai_kursus
	var existingNilaiKursus entity.NilaiKursus
	result := config.DB.Where("id_kursus = ? AND id_siswa = ? AND id_tipe_ujian = ?", 
		kursusID, siswaID, input.IdTipeUjian).First(&existingNilaiKursus)

	if result.Error == nil {
		// Record exists, update it
		existingNilaiKursus.NilaiTipeUjian = finalScore
		if err := config.DB.Save(&existingNilaiKursus).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update nilai kursus", "detail": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":           "Nilai kursus updated successfully",
			"nilai_kursus":      existingNilaiKursus,
			"total_nilai_raw":   totalNilai,
			"persentase":        persentase.Persentase,
			"nilai_final":       finalScore,
		})
	} else {
		// Record doesn't exist, create new one
		nilaiKursus := entity.NilaiKursus{
			NilaiTipeUjian: finalScore,
			IdKursus:       kursusID,
			IdSiswa:        siswaID,
			IdTipeUjian:    input.IdTipeUjian,
		}

		// Insert the new NilaiKursus into the database
		if err := config.DB.Create(&nilaiKursus).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save nilai kursus", "detail": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":           "Nilai kursus added successfully",
			"nilai_kursus":      nilaiKursus,
			"total_nilai_raw":   totalNilai,
			"persentase":        persentase.Persentase,
			"nilai_final":       finalScore,
		})
	}
}