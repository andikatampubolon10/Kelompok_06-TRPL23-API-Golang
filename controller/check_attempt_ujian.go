package controller

import (
	"cbt-api/config"
	"cbt-api/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CheckQuizAttemptRequest struct {
	IDUjian string `json:"id_ujian"`
	IDSiswa int    `json:"id_siswa"`
}

type CheckQuizAttemptResponse struct {
	HasAttempted bool   `json:"hasAttempted"`
	Message      string `json:"message,omitempty"`
}

func CheckQuizAttempt(c *gin.Context) {
	var input CheckQuizAttemptRequest

	// Try to bind JSON request body
	if err := c.ShouldBindJSON(&input); err != nil {
		// Coba dari query jika body gagal
		idUjian := c.Query("id_ujian")
		idSiswaStr := c.Query("id_siswa")

		if idUjian == "" || idSiswaStr == "" {
			fmt.Println("Missing id_ujian or id_siswa") // Debug log
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id_ujian and id_siswa must be provided",
			})
			return
		}

		idSiswa, err := strconv.Atoi(idSiswaStr)
		if err != nil {
			fmt.Println("Invalid id_siswa:", err) // Debug log
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_siswa"})
			return
		}

		input.IDUjian = idUjian
		input.IDSiswa = idSiswa
	}

	fmt.Println("Received input:", input) // Debug log

	// Konversi id_ujian dari string ke int
	idUjianInt, err := strconv.Atoi(input.IDUjian)
	if err != nil {
		fmt.Println("Invalid id_ujian:", err) // Debug log
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_ujian"})
		return
	}

	// Query menggunakan GORM
	var count int64
	if err := config.DB.Model(&entity.TipeNilai{}).
		Where("id_ujian = ? AND id_siswa = ?", idUjianInt, input.IDSiswa).
		Count(&count).Error; err != nil {
		fmt.Println("Database error:", err) // Debug log
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
		return
	}

	hasAttempted := count > 0
	message := "Student has not attempted this quiz yet"
	if hasAttempted {
		message = "Student has already attempted this quiz"
	}

	c.JSON(http.StatusOK, CheckQuizAttemptResponse{
		HasAttempted: hasAttempted,
		Message:      message,
	})
}
