package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"cbt-api/config"
	"cbt-api/entity"
)

func GetTotalNilaiByTipeUjian(c *gin.Context) {
    idKursus := c.Param("id_kursus")
    idSiswa := c.Param("id_siswa")

    var tipeNilai []entity.TipeNilai
    result := make(map[uint64]float64) // map to hold total score per tipe_ujian

    // Fetch all TipeNilai records for the given course (id_kursus) and student (id_siswa)
    err := config.DB.
        Where("tipe_nilai.id_siswa = ? AND ujian.id_kursus = ?", idSiswa, idKursus).
        Joins("JOIN ujian ON ujian.id_ujian = tipe_nilai.id_ujian").
        Joins("JOIN tipe_ujian ON tipe_ujian.id_tipe_ujian = tipe_nilai.id_tipe_ujian").
        Find(&tipeNilai).Error

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data tipe nilai", "detail": err.Error()})
        return
    }

    // Check if no data was found for the provided ids
    if len(tipeNilai) == 0 {
        c.JSON(http.StatusOK, gin.H{
            "message":             "No nilai found for the student in the given course",
            "nilai_total_per_tipe_ujian": result,
        })
        return
    }

    // Sum the nilai based on id_tipe_ujian
    for _, tn := range tipeNilai {
        result[tn.IdTipeUjian] += tn.Nilai // Add the score for each TipeUjian
    }

    // Return the total nilai per tipe_ujian
    c.JSON(http.StatusOK, gin.H{
        "message":               "Total Nilai per Tipe Ujian retrieved successfully",
        "nilai_total_per_tipe_ujian": result,
    })
}
