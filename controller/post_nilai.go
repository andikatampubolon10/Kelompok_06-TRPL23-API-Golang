package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"cbt-api/config"
	"cbt-api/entity"
	"gorm.io/gorm"
	"fmt" // Untuk debugging
	"errors"
	"strconv" // Untuk konversi string ke integer
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

	// Loop untuk mengalikan nilai_kursus dengan persentase per tipe_ujian
	var totalNilai float64
	for _, nilai := range nilaiKursus {
		// Ambil persentase berdasarkan id_tipe_ujian
		var persentase entity.Persentase
		err = config.DB.
			Where("id_kursus = ? AND id_tipe_ujian = ?", idKursus, nilai.IdTipeUjian).
			First(&persentase).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data persentase", "detail": err.Error()})
			return
		}

		// Pastikan persentase diubah menjadi desimal (misalnya 30% menjadi 0.30)
		decimalPersentase := persentase.Persentase / 100

		// Hitung nilai dengan mengalikan nilai_kursus dengan persentase
		calculatedScore := nilai.NilaiTipeUjian * decimalPersentase
		totalNilai += calculatedScore
	}

	// Debugging: Cek total nilai yang dihitung
	fmt.Println("Total Nilai setelah dihitung:", totalNilai)

	// Buat objek nilai untuk disimpan ke dalam tabel nilai
	// Pastikan id_tipe_nilai selalu 1
	nilai := entity.Nilai{
		IdKursus:  uint64(idKursus), // id_kursus dari URL parameter
		IdSiswa:   uint64(idSiswa),  // id_siswa dari URL parameter
		NilaiTotal: totalNilai,
		IdTipeNilai: 2, // Set id_tipe_nilai menjadi 1
	}

	// Simpan nilai ke dalam tabel nilai
	if err := config.DB.Create(&nilai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan nilai", "detail": err.Error()})
		return
	}

	// Kembalikan response success
	c.JSON(http.StatusOK, gin.H{
		"message": "Nilai berhasil disimpan",
		"nilai":   nilai,
	})
}
