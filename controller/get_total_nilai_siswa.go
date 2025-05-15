package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"cbt-api/config"
	"cbt-api/entity"
	"gorm.io/gorm"
)

func GetTotalNilaiSiswaByUjian(c *gin.Context) {
	idUjian := c.Param("id_ujian")
	idSiswa := c.Param("id_siswa")

	var jawabanSiswa []entity.JawabanSiswa

	err := config.DB.
    Joins("JOIN soal ON soal.id_soal = jawaban_siswa.id_soal").
    Joins("LEFT JOIN jawaban_soal ON jawaban_soal.id_jawaban_soal = jawaban_siswa.id_jawaban_soal").
    Joins("LEFT JOIN siswa ON siswa.id_siswa = jawaban_siswa.id_siswa").
    Preload("Soal").
    Preload("JawabanSoal", func(db *gorm.DB) *gorm.DB {
        return db.Select("id_jawaban_soal", "benar") // preload hanya yang perlu
    }).
    Preload("Siswa").
    Where("jawaban_siswa.id_siswa = ? AND soal.id_ujian = ?", idSiswa, idUjian).
    Find(&jawabanSiswa).Error


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data jawaban siswa", "detail": err.Error()})
		return
	}

	if len(jawabanSiswa) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":     "Jawaban Siswa retrieved successfully, but no data found",
			"nilai_total": 0,
		})
		return
	}

	var totalNilai float64
	for _, js := range jawabanSiswa {
		if js.Soal.IdSoal != 0 && js.JawabanSoal.IdJawabanSoal != 0 && js.JawabanSoal.Benar {
			totalNilai += js.Soal.NilaiPerSoal
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Total Nilai Siswa retrieved successfully",
		"nilai_total": totalNilai,
	})
}
