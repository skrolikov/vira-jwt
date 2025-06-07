package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string, tokenType string, duration time.Duration, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    tokenType,
		"exp":     time.Now().Add(duration).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}

func GenerateAccessToken(userID, secret string, duration time.Duration) (string, error) {
	return GenerateToken(userID, "access", duration, secret)
}

func GenerateRefreshToken(userID, secret string, duration time.Duration) (string, error) {
	return GenerateToken(userID, "refresh", duration, secret)
}

func IsTokenType(claims jwt.MapClaims, t string) bool {
	val, ok := claims["type"].(string)
	return ok && val == t
}

func ParseToken(tokenStr string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
