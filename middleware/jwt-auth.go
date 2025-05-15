package middleware

import (
    "github.com/gin-gonic/gin"
    "cbt-api/auth"
    "net/http"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
            c.Abort()
            return
        }

        // Verifikasi token
        token, err := auth.VerifyJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Jika token valid, lanjutkan ke handler berikutnya
        c.Set("user", token.Claims) // Menyimpan claims token jika diperlukan
        c.Next()
    }
}
