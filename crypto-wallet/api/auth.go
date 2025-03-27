package server

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your_secret_key") // Change this to a strong secret

// Claims struct for JWT
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken checks if a token is valid
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
