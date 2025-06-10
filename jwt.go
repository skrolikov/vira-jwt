package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ErrInvalidSigningMethod возвращается, если метод подписи JWT не совпадает с ожидаемым.
var ErrInvalidSigningMethod = errors.New("invalid signing method")

// GenerateToken создаёт JWT-токен с заданным userID, типом токена и временем жизни.
// Подписывается с помощью HS256 и переданного секрета.
func GenerateToken(userID string, tokenType string, duration time.Duration, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    tokenType,
		"exp":     time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GenerateAccessToken создаёт JWT-токен типа "access" с заданным временем жизни.
func GenerateAccessToken(userID, secret string, duration time.Duration) (string, error) {
	return GenerateToken(userID, "access", duration, secret)
}

// GenerateRefreshToken создаёт JWT-токен типа "refresh" с заданным временем жизни.
func GenerateRefreshToken(userID, secret string, duration time.Duration) (string, error) {
	return GenerateToken(userID, "refresh", duration, secret)
}

// IsTokenType проверяет, что в claims поле "type" совпадает с переданным значением.
func IsTokenType(claims jwt.MapClaims, t string) bool {
	val, ok := claims["type"].(string)
	return ok && val == t
}

// ParseToken парсит JWT-токен из строки и проверяет подпись с помощью секрета.
// Возвращает claims, если токен валиден, иначе ошибку.
func ParseToken(tokenStr string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Безопасность: проверяем, что метод подписи соответствует ожиданиям
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
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
