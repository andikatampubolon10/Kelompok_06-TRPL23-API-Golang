package controller

import (
	"cbt-api/config"
	"gorm.io/gorm"
	"cbt-api/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSoalByUjian(c *gin.Context) {
	// Mendapatkan parameter id_ujian dari URL
	idUjian := c.Param("id_ujian")

	// Query soal berdasarkan id_ujian
	var soal []entity.Soal
	if err := config.DB.Where("id_ujian = ?", idUjian).Find(&soal).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Soal tidak ditemukan"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Terjadi kesalahan saat mengambil data soal", "error": err.Error()})
		}
		return
	}

	// Menyusun data soal dengan jawaban yang terorganisir
	var soalDenganJawaban []map[string]interface{}

	// Untuk setiap soal, ambil jawaban yang terkait
	for _, s := range soal {
		// Query jawaban yang terkait dengan soal berdasarkan id_soal
		var jawaban []entity.JawabanSoal
		if err := config.DB.Where("id_soal = ?", s.IdSoal).Find(&jawaban).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Terjadi kesalahan saat mengambil data jawaban", "error": err.Error()})
			return
		}

		// Menyusun data soal dengan jawaban
		soalDenganJawaban = append(soalDenganJawaban, map[string]interface{}{
			"soal":    s,
			"jawaban": jawaban,
		})
	}

	// Mengirimkan respon dengan soal dan jawaban
	c.JSON(http.StatusOK, gin.H{
		"data": soalDenganJawaban,
	})
}
