package controller

import (
    "cbt-api/config"
	"cbt-api/entity"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func CalculateAndSaveScore(c *gin.Context) {
    idUjian := c.Param("id_ujian")
    idSiswa := c.Param("id_siswa")
    idTipeUjian := c.Param("id_tipe_ujian")

    // Parse params to uint64
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

    tipeUjianID, err := strconv.ParseUint(idTipeUjian, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tipe ujian ID"})
        return
    }

    // Calculate the score
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

    // Save the score to tipe_nilai table
    tipeNilai := entity.TipeNilai{
        Nilai:      totalScore,
        IdTipeUjian: tipeUjianID,
        IdSiswa:    siswaID,
        IdUjian:    ujianID,
    }

    // Check if a record already exists
    var existingTipeNilai entity.TipeNilai
    result := config.DB.Where("id_siswa = ? AND id_ujian = ? AND id_tipe_ujian = ?", 
                             siswaID, ujianID, tipeUjianID).First(&existingTipeNilai)
    
    if result.Error == nil {
        // Update existing record
        existingTipeNilai.Nilai = totalScore
        if err := config.DB.Save(&existingTipeNilai).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update score", "detail": err.Error()})
            return
        }
    } else {
        // Create new record
        if err := config.DB.Create(&tipeNilai).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save score", "detail": err.Error()})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "id_ujian":     ujianID,
        "id_siswa":     siswaID,
        "id_tipe_ujian": tipeUjianID,
        "total_score":  totalScore,
        "message":      "Score calculated and saved successfully",
    })
}
