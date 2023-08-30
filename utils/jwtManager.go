package Utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func EncodeJwtToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func DecodeJwtToken(token string) (string, error) {
	return "", nil
}
