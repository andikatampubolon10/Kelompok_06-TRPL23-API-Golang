package controller

import (
	"cbt-api/config"
	"cbt-api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLatihanSoal will return LatihanSoal details based on id_kurikulum, id_kelas, and id_mata_pelajaran
func GetLatihanSoal(c *gin.Context) {
	idKurikulum := c.Param("id_kurikulum")        // Ambil id_kurikulum dari parameter URL
	idKelas := c.Param("id_kelas")                // Ambil id_kelas dari parameter URL
	idMataPelajaran := c.Param("id_mata_pelajaran") // Ambil id_mata_pelajaran dari parameter URL

	// Cari latihan soal berdasarkan id_kurikulum, id_kelas, dan id_mata_pelajaran
	var latihanSoal []entity.Latihan
	if err := config.DB.Where("id_kurikulum = ? AND id_kelas = ? AND id_mata_pelajaran = ?", idKurikulum, idKelas, idMataPelajaran).Find(&latihanSoal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching latihan soal"})
		return
	}

	// Jika tidak ada latihan soal yang ditemukan
	if len(latihanSoal) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Latihan soal not found for the given id_kurikulum, id_kelas, and id_mata_pelajaran"})
		return
	}

	// Kirimkan data latihan soal ke client
	c.JSON(http.StatusOK, gin.H{
		"message":    "Latihan soal retrieved successfully",
		"latihan_soal": latihanSoal,
	})
}
