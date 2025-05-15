package controller

import (
    "github.com/gin-gonic/gin"
    "cbt-api/config"
    "cbt-api/entity"
    "net/http"
)

// Fungsi untuk mendapatkan mata pelajaran berdasarkan id_kurikulum
func GetMataPelajaranByKurikulum(c *gin.Context) {
    idKurikulum := c.Param("id_kurikulum") // Ambil id_kurikulum dari parameter URL

    // Cari mata pelajaran berdasarkan id_kurikulum
    var mataPelajaranList []entity.MataPelajaran
    if err := config.DB.Where("id_kurikulum = ?", idKurikulum).Find(&mataPelajaranList).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching mata pelajaran"})
        return
    }

    if len(mataPelajaranList) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "No mata pelajaran found for the given id_kurikulum"})
        return
    }

    // Kirimkan daftar mata pelajaran ke client
    c.JSON(http.StatusOK, gin.H{
        "message":        "Mata pelajaran retrieved successfully",
        "mata_pelajaran": mataPelajaranList,
    })
}
