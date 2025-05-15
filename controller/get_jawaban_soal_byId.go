package controller

import (
    "cbt-api/config"
    "cbt-api/entity"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func GetJawabanSoalByID(c *gin.Context) {
    idParam := c.Param("id_jawaban_soal")

    // Konversi ke uint64
    id, err := strconv.ParseUint(idParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var jawabanSoal entity.JawabanSoal
    if err := config.DB.First(&jawabanSoal, "id_jawaban_soal = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Jawaban soal not found", "detail": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":       "Data retrieved successfully",
        "jawaban_soal":  jawabanSoal,
    })
}
