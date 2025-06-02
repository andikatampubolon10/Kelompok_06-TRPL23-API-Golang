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
		// Record exists, update it with the new calculated score
		oldScore := existingNilaiKursus.NilaiTipeUjian
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
			"previous_score":    oldScore,
			"score_difference":  finalScore - oldScore,
			"operation":         "UPDATE",
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
			"operation":         "CREATE",
		})
	}
}

// Alternative PUT method for explicit updates
func PutNilaiKursus(c *gin.Context) {
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

	// Check if record exists first (PUT requires existing record)
	var existingNilaiKursus entity.NilaiKursus
	if err := config.DB.Where("id_kursus = ? AND id_siswa = ? AND id_tipe_ujian = ?", 
		kursusID, siswaID, input.IdTipeUjian).First(&existingNilaiKursus).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nilai kursus not found for update", "detail": err.Error()})
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

	// Step 4: Update the existing record
	oldScore := existingNilaiKursus.NilaiTipeUjian
	existingNilaiKursus.NilaiTipeUjian = finalScore
	
	if err := config.DB.Save(&existingNilaiKursus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update nilai kursus", "detail": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":           "Nilai kursus updated successfully via PUT",
		"nilai_kursus":      existingNilaiKursus,
		"total_nilai_raw":   totalNilai,
		"persentase":        persentase.Persentase,
		"nilai_final":       finalScore,
		"previous_score":    oldScore,
		"score_difference":  finalScore - oldScore,
		"operation":         "PUT_UPDATE",
	})
}

// Function to recalculate all scores for a specific student and course
func RecalculateNilaiKursus(c *gin.Context) {
	// Extract id_kursus and id_siswa from URL parameters
	idKursus := c.Param("id_kursus")
	idSiswa := c.Param("id_siswa")

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

	// Get all distinct tipe_ujian for this student and course
	var tipeUjianList []uint64
	if err := config.DB.Model(&entity.TipeNilai{}).
		Where("id_siswa = ?", siswaID).
		Distinct("id_tipe_ujian").
		Pluck("id_tipe_ujian", &tipeUjianList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tipe ujian list", "detail": err.Error()})
		return
	}

	var updatedScores []map[string]interface{}

	// Recalculate for each tipe_ujian
	for _, tipeUjianID := range tipeUjianList {
		// Check if persentase exists for this tipe_ujian and kursus
		var persentase entity.Persentase
		if err := config.DB.Where("id_tipe_ujian = ? AND id_kursus = ?", tipeUjianID, kursusID).
			First(&persentase).Error; err != nil {
			continue // Skip if no percentage found
		}

		// Calculate total nilai for this tipe_ujian
		var totalNilai float64
		if err := config.DB.Model(&entity.TipeNilai{}).
			Where("id_tipe_ujian = ? AND id_siswa = ?", tipeUjianID, siswaID).
			Select("COALESCE(SUM(nilai), 0)").
			Scan(&totalNilai).Error; err != nil {
			continue // Skip if calculation fails
		}

		// Calculate final score
		finalScore := totalNilai * (persentase.Persentase / 100)

		// Update or create nilai_kursus record
		var nilaiKursus entity.NilaiKursus
		result := config.DB.Where("id_kursus = ? AND id_siswa = ? AND id_tipe_ujian = ?", 
			kursusID, siswaID, tipeUjianID).First(&nilaiKursus)

		if result.Error == nil {
			// Update existing record
			nilaiKursus.NilaiTipeUjian = finalScore
			config.DB.Save(&nilaiKursus)
		} else {
			// Create new record
			nilaiKursus = entity.NilaiKursus{
				NilaiTipeUjian: finalScore,
				IdKursus:       kursusID,
				IdSiswa:        siswaID,
				IdTipeUjian:    tipeUjianID,
			}
			config.DB.Create(&nilaiKursus)
		}

		updatedScores = append(updatedScores, map[string]interface{}{
			"id_tipe_ujian":   tipeUjianID,
			"total_nilai_raw": totalNilai,
			"persentase":      persentase.Persentase,
			"nilai_final":     finalScore,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "All scores recalculated successfully",
		"updated_scores": updatedScores,
		"total_updated":  len(updatedScores),
	})
}