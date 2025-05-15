package controller

import (
    "github.com/gin-gonic/gin"
    "cbt-api/config"
    "cbt-api/entity"
    "cbt-api/auth"
    "golang.org/x/crypto/bcrypt"
    "net/http"
)

func Login(c *gin.Context) {
    var userInput struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&userInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // 🔍 Cari user berdasarkan email
    var user entity.Users
    if err := config.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    // ✅ Verifikasi password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
        return
    }

    // 🔗 Ambil siswa yang terkait dengan user ini
    var siswa entity.Siswa
    if err := config.DB.Where("id_user = ?", user.Id).First(&siswa).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Siswa tidak ditemukan untuk user ini"})
        return
    }

    // 🔐 Buat token JWT
    token, err := auth.GenerateJWT(user.Id, user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    // 🎯 Kirimkan token dan id_siswa
    c.JSON(http.StatusOK, gin.H{
        "message":   "Login successful",
        "token":     token,
        "id_siswa":  siswa.IdSiswa,     // Inilah yang kamu cari
        "nama_siswa":   siswa.NamaSiswa,   // Optional
    })
}
