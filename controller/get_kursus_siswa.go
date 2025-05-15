package controller

import (
    "github.com/gin-gonic/gin"
    "cbt-api/config"
    "cbt-api/entity"
    "net/http"
)

func GetKursusBySiswa(c *gin.Context) {
    idSiswa := c.Param("id_siswa")

    var kursusSiswa []entity.KursusSiswa

    if err := config.DB.Preload("Kursus").Where("id_siswa = ?", idSiswa).Find(&kursusSiswa).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data kursus siswa"})
        return
    }

    var kursusList []gin.H
    for _, ks := range kursusSiswa {
        // Pastikan Kursus tidak kosong (jaga-jaga kalau relasi gagal)
        if ks.Kursus.IdKursus != 0 {
            kursusList = append(kursusList, gin.H{
                "id_kursus":   ks.Kursus.IdKursus,
                "nama_kursus": ks.Kursus.NamaKursus,
                "password":    ks.Kursus.Password,
                "image":       ks.Kursus.Image,
                "image_url":   ks.Kursus.ImageUrl,
            })
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Kursus retrieved successfully",
        "kursus":  kursusList,
    })
}
