package auth

import (
	"errors"
	"log"
	"time"

	"employeejwt/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userID int, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWTSecret()))
	if err != nil {
		log.Printf("auth.GenerateToken: failed for userID %d email %s: %v", userID, email, err)
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ClaimsValid := token.Method.(*jwt.SigningMethodHMAC); !ClaimsValid {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.JWTSecret()), nil
	})
	if err != nil {
		log.Printf("auth.ValidateToken: parse failed: %v", err)
		return nil, err
	}

	claims, ClaimsValid := token.Claims.(*Claims)
	if !ClaimsValid || !token.Valid {
		log.Printf("auth.ValidateToken: invalid token claims for token %q", tokenString)
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
