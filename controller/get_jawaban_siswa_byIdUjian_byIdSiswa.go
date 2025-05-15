package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"cbt-api/config"
	"cbt-api/entity"
)

func GetJawabanSiswa(c *gin.Context) {
	idUjian := c.Param("id_ujian")
	idSiswa := c.Param("id_siswa")

	var jawabanSiswa []entity.JawabanSiswa

	err := config.DB.
		Joins("JOIN soal ON soal.id_soal = jawaban_siswa.id_soal").
		Joins("LEFT JOIN jawaban_soal ON jawaban_soal.id_jawaban_soal = jawaban_siswa.id_jawaban_soal").
		Joins("LEFT JOIN siswa ON siswa.id_siswa = jawaban_siswa.id_siswa").
		Preload("Soal").
		Preload("JawabanSoal").
		Preload("Siswa").
		Where("jawaban_siswa.id_siswa = ? AND soal.id_ujian = ?", idSiswa, idUjian).
		Find(&jawabanSiswa).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data jawaban siswa", "detail": err.Error()})
		return
	}

	if len(jawabanSiswa) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Jawaban Siswa retrieved successfully, but no data found",
			"jawaban": nil,
		})
		return
	}

	var jawabanList []gin.H
	for _, js := range jawabanSiswa {
		jawabanList = append(jawabanList, gin.H{
			"id_jawaban_siswa": js.IdJawabanSiswa,
			"jawaban_siswa":    js.JawabanText,
			"benar":            js.JawabanSoal.Benar,
			"nilai_per_soal":   js.Soal.NilaiPerSoal,
			"id_soal":          js.IdSoal,
			"id_siswa":         js.IdSiswa,
			"id_ujian":         js.Soal.IdUjian,
			"grade":            js.JawabanSoal.Benar,
			"id_jawaban_soal":  js.IdJawabanSoal,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Jawaban Siswa retrieved successfully",
		"jawaban": jawabanList,
	})
}
