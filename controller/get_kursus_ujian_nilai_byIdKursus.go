package controller

import (
    "cbt-api/config"
    "cbt-api/entity"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

func GetKursusWithUjianAndNilai(c *gin.Context) {
    idKursusParam := c.Param("id_kursus")
    idSiswaParam := c.Param("id_siswa")

    idKursus, err := strconv.ParseUint(idKursusParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_kursus"})
        return
    }

    idSiswa, err := strconv.ParseUint(idSiswaParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_siswa"})
        return
    }

    // Ambil data kursus
    var kursus entity.Kursus
    if err := config.DB.First(&kursus, "id_kursus = ?", idKursus).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Kursus tidak ditemukan"})
        return
    }

    // Ambil daftar ujian terkait kursus
    var ujianList []entity.Ujian
    if err := config.DB.Where("id_kursus = ?", idKursus).
        Preload("TipeUjian").
        Find(&ujianList).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil ujian"})
        return
    }

    // Ambil nilai siswa untuk ujian dalam kursus ini
    var nilaiList []entity.TipeNilai
    if err := config.DB.Joins("JOIN ujian ON ujian.id_ujian = tipe_nilai.id_ujian").
        Where("ujian.id_kursus = ? AND tipe_nilai.id_siswa = ?", idKursus, idSiswa).
        Preload("Siswa").
        Preload("TipeUjian").
        Preload("Ujian").
        Find(&nilaiList).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil nilai siswa"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "kursus":     kursus,
        "ujian_list": ujianList,
        "nilai_list": nilaiList,
    })
}
