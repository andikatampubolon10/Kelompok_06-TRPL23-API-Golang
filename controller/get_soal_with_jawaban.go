package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"gorm.io/gorm"
	"cbt-api/config"
	"cbt-api/entity"
)

func GetSoalWithJawaban(c *gin.Context) {
	// Mendapatkan parameter id_soal dari URL
	idSoalStr := c.Param("id_soal")
	idSoal, err := strconv.ParseUint(idSoalStr, 10, 64)  // Mengubah string ke uint64
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id_soal"})
		return
	}

	// Query soal berdasarkan id_soal
	var soal entity.Soal
	if err := config.DB.Where("id_soal = ?", idSoal).First(&soal).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Soal tidak ditemukan"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Terjadi kesalahan saat mengambil data soal", "error": err.Error()})
		}
		return
	}

	// Query jawaban yang terkait dengan soal berdasarkan id_soal
	var jawaban []entity.JawabanSoal
	if err := config.DB.Where("id_soal = ?", idSoal).Find(&jawaban).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Terjadi kesalahan saat mengambil data jawaban", "error": err.Error()})
		return
	}

	// Menyusun respon dengan data soal dan jawaban
	c.JSON(http.StatusOK, gin.H{
		"soal":    soal,
		"jawaban": jawaban,
	})
}


