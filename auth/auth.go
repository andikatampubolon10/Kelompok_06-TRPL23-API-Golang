package auth

import (
    "github.com/dgrijalva/jwt-go"
    "time"
   
)

// Secret key untuk signing JWT (jaga kerahasiaannya)
var secretKey = []byte("your_secret_key")

// Fungsi untuk menghasilkan JWT
func GenerateJWT(userId uint64, email string) (string, error) {
    claims := jwt.MapClaims{
        "id":    userId,
        "email": email,
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 1 hari
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Menandatangani token
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

// Fungsi untuk memverifikasi JWT dan mengembalikan data claims
func VerifyJWT(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Pastikan token menggunakan signing method yang benar
        
        return secretKey, nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }

    return token, nil
}
