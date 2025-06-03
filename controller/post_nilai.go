package controller

import (
	"cbt-api/config"
	"cbt-api/entity"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostNilai(c *gin.Context) {
	// Ambil id_kursus dan id_siswa dari URL parameter
	idKursusStr := c.Param("id_kursus")
	idSiswaStr := c.Param("id_siswa")

	// Konversi id_kursus dan id_siswa dari string ke uint64
	idKursus, err := strconv.ParseUint(idKursusStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_kursus, must be a valid number", "detail": err.Error()})
		return
	}

	idSiswa, err := strconv.ParseUint(idSiswaStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_siswa, must be a valid number", "detail": err.Error()})
		return
	}

	// Validasi apakah id_siswa ada di tabel siswa
	var siswa entity.Siswa
	err = config.DB.Where("id_siswa = ?", idSiswa).First(&siswa).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Siswa tidak ditemukan", "id_siswa": idSiswa})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data siswa", "detail": err.Error()})
		return
	}

	// Query untuk mendapatkan semua data nilai_kursus berdasarkan id_kursus dan id_siswa
	var nilaiKursus []entity.NilaiKursus
	err = config.DB.
		Where("id_kursus = ? AND id_siswa = ?", idKursus, idSiswa).
		Find(&nilaiKursus).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data nilai kursus", "detail": err.Error()})
		return
	}

	// Cek apakah data nilai_kursus ditemukan
	if len(nilaiKursus) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tidak ada nilai untuk siswa pada kursus ini"})
		return
	}

	// Loop untuk menghitung total nilai_kursus
	var totalNilai float64
	for _, nilai := range nilaiKursus {
		totalNilai += nilai.NilaiTipeUjian
	}

	// Debugging: Cek total nilai yang dihitung
	fmt.Println("Total Nilai setelah dihitung:", totalNilai)

	// Cek apakah sudah ada record nilai untuk kursus dan siswa ini
	var existingNilai entity.Nilai
	result := config.DB.Where("id_kursus = ? AND id_siswa = ?", idKursus, idSiswa).First(&existingNilai)

	now := time.Now()

	if result.Error == nil {
		// Record sudah ada, update nilai_total
		oldTotal := existingNilai.NilaiTotal
		existingNilai.NilaiTotal = totalNilai
		existingNilai.UpdatedAt = now

		// Update record yang sudah ada
		if err := config.DB.Save(&existingNilai).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate nilai", "detail": err.Error()})
			return
		}

		// Kembalikan response success untuk update
		c.JSON(http.StatusOK, gin.H{
			"message":         "Nilai berhasil diupdate",
			"nilai":           existingNilai,
			"previous_total":  oldTotal,
			"new_total":       totalNilai,
			"total_difference": totalNilai - oldTotal,
			"operation":       "UPDATE",
		})
	} else {
		// Record belum ada, buat baru
		nilai := entity.Nilai{
			IdKursus:   uint64(idKursus),
			IdSiswa:    uint64(idSiswa),
			NilaiTotal: totalNilai,
			CreatedAt:  now,
			UpdatedAt:  now,
		}

		// Simpan nilai baru ke dalam tabel nilai
		if err := config.DB.Create(&nilai).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan nilai", "detail": err.Error()})
			return
		}

		// Kembalikan response success untuk create
		c.JSON(http.StatusOK, gin.H{
			"message":   "Nilai berhasil disimpan",
			"nilai":     nilai,
			"operation": "CREATE",
		})
	}
}

// PUT method untuk explicit update
func PutNilai(c *gin.Context) {
	// Ambil id_kursus dan id_siswa dari URL parameter
	idKursusStr := c.Param("id_kursus")
	idSiswaStr := c.Param("id_siswa")

	// Konversi id_kursus dan id_siswa dari string ke uint64
	idKursus, err := strconv.ParseUint(idKursusStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_kursus, must be a valid number", "detail": err.Error()})
		return
	}

	idSiswa, err := strconv.ParseUint(idSiswaStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_siswa, must be a valid number", "detail": err.Error()})
		return
	}

	// Cek apakah record nilai sudah ada (PUT memerlukan record yang sudah ada)
	var existingNilai entity.Nilai
	if err := config.DB.Where("id_kursus = ? AND id_siswa = ?", idKursus, idSiswa).First(&existingNilai).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Nilai tidak ditemukan untuk diupdate", "id_kursus": idKursus, "id_siswa": idSiswa})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data nilai", "detail": err.Error()})
		return
	}

	// Query untuk mendapatkan semua data nilai_kursus berdasarkan id_kursus dan id_siswa
	var nilaiKursus []entity.NilaiKursus
	err = config.DB.
		Where("id_kursus = ? AND id_siswa = ?", idKursus, idSiswa).
		Find(&nilaiKursus).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data nilai kursus", "detail": err.Error()})
		return
	}

	// Cek apakah data nilai_kursus ditemukan
	if len(nilaiKursus) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tidak ada nilai untuk siswa pada kursus ini"})
		return
	}

	// Loop untuk menghitung total nilai_kursus
	var totalNilai float64
	for _, nilai := range nilaiKursus {
		totalNilai += nilai.NilaiTipeUjian
	}

	// Update nilai_total
	oldTotal := existingNilai.NilaiTotal
	existingNilai.NilaiTotal = totalNilai
	existingNilai.UpdatedAt = time.Now()

	// Update record
	if err := config.DB.Save(&existingNilai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate nilai", "detail": err.Error()})
		return
	}

	// Kembalikan response success
	c.JSON(http.StatusOK, gin.H{
		"message":          "Nilai berhasil diupdate via PUT",
		"nilai":            existingNilai,
		"previous_total":   oldTotal,
		"new_total":        totalNilai,
		"total_difference": totalNilai - oldTotal,
		"operation":        "PUT_UPDATE",
	})
}

// Function untuk recalculate semua nilai untuk semua siswa dalam kursus tertentu
func RecalculateNilai(c *gin.Context) {
	idKursusStr := c.Param("id_kursus")

	// Konversi id_kursus dari string ke uint64
	idKursus, err := strconv.ParseUint(idKursusStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_kursus, must be a valid number", "detail": err.Error()})
		return
	}

	// Ambil semua siswa yang memiliki nilai_kursus untuk kursus ini
	var distinctSiswa []uint64
	if err := config.DB.Model(&entity.NilaiKursus{}).
		Where("id_kursus = ?", idKursus).
		Distinct("id_siswa").
		Pluck("id_siswa", &distinctSiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar siswa", "detail": err.Error()})
		return
	}

	var updatedCount int
	var results []map[string]interface{}

	// Recalculate untuk setiap siswa
	for _, siswaID := range distinctSiswa {
		// Query untuk mendapatkan semua data nilai_kursus berdasarkan id_kursus dan id_siswa
		var nilaiKursus []entity.NilaiKursus
		err = config.DB.
			Where("id_kursus = ? AND id_siswa = ?", idKursus, siswaID).
			Find(&nilaiKursus).Error

		if err != nil {
			continue // Skip jika error
		}

		if len(nilaiKursus) == 0 {
			continue // Skip jika tidak ada data
		}

		// Hitung total nilai
		var totalNilai float64
		for _, nilai := range nilaiKursus {
			totalNilai += nilai.NilaiTipeUjian
		}

		// Update atau create record nilai
		var existingNilai entity.Nilai
		result := config.DB.Where("id_kursus = ? AND id_siswa = ?", idKursus, siswaID).First(&existingNilai)

		now := time.Now()

		if result.Error == nil {
			// Update existing record
			oldTotal := existingNilai.NilaiTotal
			existingNilai.NilaiTotal = totalNilai
			existingNilai.UpdatedAt = now
			config.DB.Save(&existingNilai)

			results = append(results, map[string]interface{}{
				"id_siswa":         siswaID,
				"previous_total":   oldTotal,
				"new_total":        totalNilai,
				"total_difference": totalNilai - oldTotal,
				"operation":        "UPDATED",
			})
		} else {
			// Create new record
			nilai := entity.Nilai{
				IdKursus:   idKursus,
				IdSiswa:    siswaID,
				NilaiTotal: totalNilai,
				CreatedAt:  now,
				UpdatedAt:  now,
			}
			config.DB.Create(&nilai)

			results = append(results, map[string]interface{}{
				"id_siswa":   siswaID,
				"new_total":  totalNilai,
				"operation":  "CREATED",
			})
		}

		updatedCount++
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Recalculation completed successfully",
		"updated_count": updatedCount,
		"results":       results,
	})
}