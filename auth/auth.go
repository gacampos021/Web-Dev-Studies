package auth

import (
	"fmt"
	"study-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID, email string) (string, error) {
	config.LoadEnv()
	var jwtSecret = config.GetEnv("URI_MONGO")

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", fmt.Errorf("erro ao assinar token: %v", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	config.LoadEnv()
	var jwtSecret = config.GetEnv("URI_MONGO")

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("erro ao fazer parse do token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	return claims, nil
}
