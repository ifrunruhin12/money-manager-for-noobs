package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenIssuer = "money-manager"
	leeway      = 30 * time.Second
)

type Claims struct {
	jwt.RegisteredClaims
}

func GenerateToken(userID string, secret string, expiry time.Duration) (string, error) {
	if secret == "" {
		return "", fmt.Errorf("jwt secret is empty")
	}

	if expiry <= 0 {
		return "", fmt.Errorf("jwt expiry must be greater than zero")
	}

	now := time.Now().UTC()

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(now.Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    tokenIssuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return signed, nil
}

func ParseToken(tokenString string, secret string) (string, error) {
	if secret == "" {
		return "", fmt.Errorf("jwt secret is empty")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(t *jwt.Token) (any, error) {
			// Prevent algorithm substitution attacks
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithIssuer(tokenIssuer),
		jwt.WithLeeway(leeway),
	)

	if err != nil {
		return "", fmt.Errorf("parse token: %w", err)
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	if claims.Subject == "" {
		return "", fmt.Errorf("missing subject claim")
	}

	return claims.Subject, nil
}
