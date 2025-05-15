package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"cbt-api/config"
	"cbt-api/entity"
)

func GetAllKursusSiswa(c *gin.Context) {
	var kursusSiswa []entity.KursusSiswa
	if err := config.DB.Preload("Kursus").Preload("Siswa").Find(&kursusSiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kursusSiswa)
}

func GetKursusSiswaByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id_kursus_siswa"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var kursusSiswa entity.KursusSiswa
	if err := config.DB.Preload("Kursus").Preload("Siswa").First(&kursusSiswa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "KursusSiswa not found"})
		return
	}
	c.JSON(http.StatusOK, kursusSiswa)
}

func CreateKursusSiswa(c *gin.Context) {
	var kursusSiswa entity.KursusSiswa
	if err := c.ShouldBindJSON(&kursusSiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&kursusSiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, kursusSiswa)
}

func UpdateKursusSiswa(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id_kursus_siswa"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var kursusSiswa entity.KursusSiswa
	if err := config.DB.First(&kursusSiswa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "KursusSiswa not found"})
		return
	}

	if err := c.ShouldBindJSON(&kursusSiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&kursusSiswa).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kursusSiswa)
}

func DeleteKursusSiswa(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id_kursus_siswa"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := config.DB.Delete(&entity.KursusSiswa{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "KursusSiswa deleted successfully"})
}