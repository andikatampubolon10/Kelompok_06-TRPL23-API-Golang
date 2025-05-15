package controller

import (
	"cbt-api/config"   // Import DB configuration
	"cbt-api/entity"   // Import Kursus and KursusSiswa entities
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetAvailableKursusForSiswa will return courses that the student is not enrolled in.
// GetAvailableKursusForSiswa will return courses that the student is not enrolled in.
func GetAvailableKursusForSiswa(c *gin.Context) {
    idSiswa := c.Param("id_siswa")

    var kursusList []entity.Kursus

    // Ambil kursus yang belum diambil siswa
    err := config.DB.Table("kursus").
        Where("kursus.id_kursus NOT IN (?)",
            config.DB.Table("kursus_siswa").
                Select("id_kursus").
                Where("id_siswa = ?", idSiswa),
        ).Find(&kursusList).Error

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Unable to retrieve available courses",
            "details": err.Error(),
        })
        return
    }

    // Format output sesuai kebutuhan Flutter
    var result []gin.H
    for _, k := range kursusList {
        result = append(result, gin.H{
            "id_kursus": k.IdKursus,
            "image":     k.Image,
            "image_url": k.ImageUrl,
            "title":     k.NamaKursus,
        })
    }

    // Return response langsung sebagai array (tanpa nested "kursus")
    c.JSON(http.StatusOK, result)
}
