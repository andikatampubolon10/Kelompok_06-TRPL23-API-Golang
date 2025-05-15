// service/jwt_service.go
package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4" // Using the newer version of the JWT package
)

// JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetUserIDByToken(token string) (string, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService creates a new instance of JWTService
func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "cbt-api",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "cbt-api-secret-key" // Default secret key
	}
	return secretKey
}

func (j *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaim{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) GetUserIDByToken(token string) (string, error) {
	t, err := j.ValidateToken(token)
	if err != nil {
		return "", err
	}
	claims := t.Claims.(jwt.MapClaims)
	return claims["user_id"].(string), nil
}