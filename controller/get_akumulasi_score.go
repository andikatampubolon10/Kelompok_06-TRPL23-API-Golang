package controller

import (
    "cbt-api/config"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func CalculateScore(c *gin.Context) {
    idUjian := c.Param("id_ujian")
    idSiswa := c.Param("id_siswa")

    // Parse param ke uint64
    ujianID, err := strconv.ParseUint(idUjian, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ujian ID"})
        return
    }

    siswaID, err := strconv.ParseUint(idSiswa, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid siswa ID"})
        return
    }

    // Query SQL langsung untuk hitung total nilai
    var totalScore float64
    query := `
        SELECT COALESCE(SUM(s.nilai_per_soal), 0) as total_score
        FROM jawaban_siswa js
        JOIN jawaban_soal j ON js.id_jawaban_soal = j.id_jawaban_soal
        JOIN soal s ON js.id_soal = s.id_soal
        WHERE js.id_siswa = ?
        AND s.id_ujian = ?
        AND j.benar = true
    `

    if err := config.DB.Raw(query, siswaID, ujianID).Scan(&totalScore).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate score", "detail": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id_ujian":     ujianID,
        "id_siswa":     siswaID,
        "total_score":  totalScore,
    })
}
