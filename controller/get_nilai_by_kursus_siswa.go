package controller

import (
	"github.com/gin-gonic/gin"
	"cbt-api/config"
	"cbt-api/entity"
	"net/http"
	"strconv"
)

func GetNilaiByKursusAndSiswa(c *gin.Context) {
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

	// Query untuk mendapatkan data nilai berdasarkan id_kursus dan id_siswa
	var nilai []entity.Nilai
	err = config.DB.Preload("Kursus").Preload("Siswa").
		Where("id_kursus = ? AND id_siswa = ?", idKursus, idSiswa).
		Find(&nilai).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data nilai", "detail": err.Error()})
		return
	}

	// Jika tidak ada data yang ditemukan
	if len(nilai) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Data nilai tidak ditemukan untuk kursus dan siswa ini"})
		return
	}

	// Hitung total nilai dari semua record yang ditemukan
	var totalNilai float64 = 0
	for _, n := range nilai {
		totalNilai += n.NilaiTotal
	}

	// Siapkan response data dengan detail nilai dan total
	var nilaiDetails []map[string]interface{}
	for _, n := range nilai {
		nilaiDetail := map[string]interface{}{
			"id_nilai":     n.IdNilai,
			"nilai_total":  n.NilaiTotal,
			"id_kursus":    n.IdKursus,
			"id_siswa":     n.IdSiswa,
			"created_at":   n.CreatedAt,
			"updated_at":   n.UpdatedAt,
			"kursus":       n.Kursus,
			"siswa":        n.Siswa,
		}
		nilaiDetails = append(nilaiDetails, nilaiDetail)
	}

	// Kembalikan response dengan data nilai yang ditemukan dan total nilai
	c.JSON(http.StatusOK, gin.H{
		"message":     "Data nilai berhasil diambil",
		"nilai":       nilaiDetails,
		"total_nilai": totalNilai,
		"jumlah_data": len(nilai),
		"summary": map[string]interface{}{
			"id_kursus":   idKursus,
			"id_siswa":    idSiswa,
			"total_nilai": totalNilai,
			"rata_rata":   totalNilai / float64(len(nilai)),
		},
	})
}

// Alternative function specifically for getting only the total score
func GetTotalNilaiByKursusAndSiswa(c *gin.Context) {
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

	// Query untuk menghitung total nilai menggunakan SUM
	var totalNilai float64
	err = config.DB.Model(&entity.Nilai{}).
		Where("id_kursus = ? AND id_siswa = ?", idKursus, idSiswa).
		Select("COALESCE(SUM(nilai_total), 0)").
		Scan(&totalNilai).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil total nilai", "detail": err.Error()})
		return
	}

	// Query untuk menghitung jumlah record
	var jumlahData int64
	err = config.DB.Model(&entity.Nilai{}).
		Where("id_kursus = ? AND id_siswa = ?", idKursus, idSiswa).
		Count(&jumlahData).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung jumlah data", "detail": err.Error()})
		return
	}

	// Jika tidak ada data yang ditemukan
	if jumlahData == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":     "Data nilai tidak ditemukan untuk kursus dan siswa ini",
			"total_nilai": 0,
			"jumlah_data": 0,
			"rata_rata":   0,
		})
		return
	}

	// Hitung rata-rata
	rataRata := totalNilai / float64(jumlahData)

	// Kembalikan response dengan total nilai
	c.JSON(http.StatusOK, gin.H{
		"message":     "Total nilai berhasil dihitung",
		"total_nilai": totalNilai,
		"jumlah_data": jumlahData,
		"rata_rata":   rataRata,
		"summary": map[string]interface{}{
			"id_kursus":   idKursus,
			"id_siswa":    idSiswa,
			"total_nilai": totalNilai,
			"rata_rata":   rataRata,
		},
	})
}