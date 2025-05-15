package controller

import (
	"gorm.io/gorm"
	"cbt-api/config"
	"cbt-api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllSiswa mengembalikan seluruh data siswa beserta data kelas yang terkait
// GetSiswaWithKelas mengembalikan data siswa beserta kelasnya berdasarkan id_siswa
func GetSiswaWithKelas(c *gin.Context) {
	// Ambil id_siswa dari parameter URL
	id := c.Param("id_siswa")

	// Inisialisasi model Siswa
	var siswa entity.Siswa

	// Query untuk mengambil data siswa beserta relasi kelas
	err := config.DB.
		Preload("Kelas"). // Memastikan relasi Kelas dimuat
		First(&siswa, "id_siswa = ?", id).Error

	// Jika terjadi error
	if err != nil {
		// Jika error karena data tidak ditemukan
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Siswa tidak ditemukan"})
			return
		}

		// Jika terjadi error lain
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan pada server"})
		return
	}

	// Membuat response untuk nama_siswa, nis, nama_kelas, dan id_siswa
	response := gin.H{
		"id_siswa":   siswa.IdSiswa,    // Menambahkan id_siswa ke response
		"nama_siswa": siswa.NamaSiswa,
		"nis":        siswa.NIS,
		"nama_kelas": siswa.Kelas.NamaKelas, // Pastikan Kelas sudah preload dengan benar
	}

	// Kirimkan response dalam format JSON
	c.JSON(http.StatusOK, response)
}
