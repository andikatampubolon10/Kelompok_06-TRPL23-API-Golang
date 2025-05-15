package controller

import (
	"cbt-api/config"
	"cbt-api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllKurikulum will return all the curriculum records.
func GetAllKurikulum(c *gin.Context) {
	var kurikulumList []entity.Kurikulum

	// Retrieve all curriculum from the database
	err := config.DB.Find(&kurikulumList).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve curriculum"})
		return
	}

	// Return the list of curriculum in JSON format
	c.JSON(http.StatusOK, gin.H{
		"message":  "Kurikulum retrieved successfully",
		"kurikulum": kurikulumList,
	})
}
