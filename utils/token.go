package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)


func GenerateAccessToken(userID uint) (string, error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": userID,
	"exp": time.Now().Add(15 * time.Minute).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func GenerateRefreshToken() string {
	return uuid.NewString()
}