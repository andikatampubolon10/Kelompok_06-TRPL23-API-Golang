package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"cbt-api/config"
	"cbt-api/entity"
)

func CreateTipeNilai(c *gin.Context) {
	// Bind the JSON body to the TipeNilai struct
	var tipeNilai entity.TipeNilai
	if err := c.ShouldBindJSON(&tipeNilai); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Save the new TipeNilai to the database
	if err := config.DB.Create(&tipeNilai).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save score"})
		return
	}

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Score submitted successfully",
		"data":    tipeNilai,
	})
}
