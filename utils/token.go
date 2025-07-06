package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)


func GenerateAccessToken(userID uint) (string, error){
	claims := jwt.MapClaims{
		"user_id":userID,
		"exp":time.Now().Add(15 * time.Minute).Unix(),
	}
	 token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	 return token.SignedString([]byte(os.Getenv("SECRET")))
}

func GenerateRefreshToken() string {
	return uuid.NewString()
}