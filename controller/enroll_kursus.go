package controller

import (
    "cbt-api/config"
    "cbt-api/entity"
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
)

func EnrollKursusSiswa(c *gin.Context) {
    var input struct {
        IdSiswa  uint64 `json:"id_siswa"`  // Menggunakan uint64 sesuai dengan entity KursusSiswa
        IdKursus uint64 `json:"id_kursus"` // Menggunakan uint64 sesuai dengan entity KursusSiswa
    }

    // Bind JSON ke input struct
    if err := c.ShouldBindJSON(&input); err != nil {
        fmt.Println("Error binding JSON:", err) // Debugging log
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
        return
    }

    // Debugging: Print input yang diterima
    fmt.Println("Received input:", input)

    // Pastikan input IdSiswa dan IdKursus valid
    if input.IdSiswa == 0 || input.IdKursus == 0 {
        fmt.Println("Missing id_siswa or id_kursus") // Debugging log
        c.JSON(http.StatusBadRequest, gin.H{"error": "id_siswa and id_kursus must be provided"})
        return
    }

    // Mencari kursus berdasarkan id_kursus
    var kursus entity.Kursus
    if err := config.DB.Where("id_kursus = ?", input.IdKursus).First(&kursus).Error; err != nil {
        fmt.Println("Kursus not found:", err) // Debugging log
        c.JSON(http.StatusNotFound, gin.H{"error": "Kursus not found"})
        return
    }

    // Mencari siswa berdasarkan id_siswa
    var siswa entity.Siswa
    if err := config.DB.Where("id_siswa = ?", input.IdSiswa).First(&siswa).Error; err != nil {
        fmt.Println("Siswa not found:", err) // Debugging log
        c.JSON(http.StatusNotFound, gin.H{"error": "Siswa not found"})
        return
    }

    // Membuat entri baru di KursusSiswa
    kursusSiswa := entity.KursusSiswa{
        IdSiswa:  input.IdSiswa,  // Menggunakan uint64
        IdKursus: input.IdKursus, // Menggunakan uint64
    }

    if err := config.DB.Create(&kursusSiswa).Error; err != nil {
        fmt.Println("Error creating KursusSiswa:", err) // Debugging log
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll in course"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Successfully enrolled in the course",
        "kursus": kursus.NamaKursus,
    })
}
