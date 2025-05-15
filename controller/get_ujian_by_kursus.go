package controller

import (
    "github.com/gin-gonic/gin"
    "cbt-api/config"
    "cbt-api/entity"
    "net/http"
)

// Fungsi untuk mendapatkan ujian dan materi berdasarkan kursus
func GetUjianAndMateriByKursus(c *gin.Context) {
    idKursus := c.Param("id_kursus") // Ambil id_kursus dari parameter URL

    // Cari kursus berdasarkan id_kursus
    var kursus entity.Kursus
    if err := config.DB.Where("id_kursus = ?", idKursus).First(&kursus).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Kursus not found"})
        return
    }

    // Ambil semua ujian yang terkait dengan kursus tersebut
    var ujianList []entity.Ujian
    if err := config.DB.Where("id_kursus = ?", kursus.IdKursus).Find(&ujianList).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching ujian"})
        return
    }

    // Ambil semua materi yang terkait dengan kursus tersebut
    var materiList []entity.Materi
    if err := config.DB.Where("id_kursus = ?", kursus.IdKursus).Find(&materiList).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching materi"})
        return
    }

    // Kirimkan daftar ujian, materi, dan nama kursus ke client
    c.JSON(http.StatusOK, gin.H{
        "message":    "Ujian and Materi retrieved successfully",
        "kursus_name": kursus.NamaKursus, // Kirimkan nama kursus
        "ujian":      ujianList,
        "materi":     materiList, // Kirimkan materi
    })
}