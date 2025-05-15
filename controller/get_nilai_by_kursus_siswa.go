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
	err = config.DB.Preload("Kursus").Preload("Siswa").Preload("TipeNilai").
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

	// Kembalikan response dengan data nilai yang ditemukan
	c.JSON(http.StatusOK, gin.H{
		"message": "Data nilai berhasil diambil",
		"nilai":   nilai,
	})
}
