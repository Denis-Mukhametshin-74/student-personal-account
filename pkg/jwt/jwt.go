package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	StudentID int    `json:"student_id"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

type Config struct {
	SecretKey  string
	Expiration string
}

var (
	jwtSecret     []byte
	jwtExpiration time.Duration
)

func InitJWT(cfg Config) error {
	if cfg.SecretKey == "" {
		return errors.New("JWT_SECRET is required")
	}
	jwtSecret = []byte(cfg.SecretKey)

	expirationStr := cfg.Expiration
	if expirationStr == "" {
		expirationStr = "24h"
	}

	var err error
	jwtExpiration, err = time.ParseDuration(expirationStr)
	if err != nil {
		if hours, err := strconv.Atoi(expirationStr); err == nil {
			jwtExpiration = time.Duration(hours) * time.Hour
		} else {
			return errors.New("invalid JWT_EXPIRATION format")
		}
	}

	return nil
}

func GenerateToken(studentID int, email string) (string, error) {
	expirationTime := time.Now().Add(jwtExpiration)

	claims := &Claims{
		StudentID: studentID,
		Email:     email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
