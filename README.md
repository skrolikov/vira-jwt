# Vira JWT

Пакет `vira-jwt` предоставляет удобные функции для работы с JWT токенами, включая генерацию, парсинг и проверку токенов.

## Особенности

- Генерация access и refresh токенов
- Проверка подписи токенов
- Поддержка JWT с алгоритмом HS256
- Проверка типа токена (access/refresh)
- Простая интеграция с Go-приложениями

## Установка

```bash
go get github.com/skrolikov/vira-jwt
```

## Использование

### Генерация токенов

```go
import "github.com/skrolikov/vira-jwt"

func main() {
    secret := "your-secret-key"
    userID := "user123"
    
    // Генерация access токена (короткое время жизни)
    accessToken, err := jwt.GenerateAccessToken(userID, secret, 15*time.Minute)
    if err != nil {
        // Обработка ошибки
    }
    
    // Генерация refresh токена (долгое время жизни)
    refreshToken, err := jwt.GenerateRefreshToken(userID, secret, 7*24*time.Hour)
    if err != nil {
        // Обработка ошибки
    }
}
```

### Парсинг и проверка токенов

```go
import "github.com/skrolikov/vira-jwt"

func main() {
    tokenString := "your.jwt.token.here"
    secret := "your-secret-key"
    
    claims, err := jwt.ParseToken(tokenString, secret)
    if err != nil {
        // Токен невалиден
        return
    }
    
    // Проверка типа токена
    if jwt.IsTokenType(claims, "access") {
        // Это access токен
    } else if jwt.IsTokenType(claims, "refresh") {
        // Это refresh токен
    }
    
    // Получение userID из claims
    userID := claims["user_id"].(string)
}
```

## Функции

### Основные функции

- `GenerateToken(userID, tokenType, duration, secret)` - генерирует JWT токен с указанными параметрами
- `GenerateAccessToken(userID, secret, duration)` - генерирует access токен
- `GenerateRefreshToken(userID, secret, duration)` - генерирует refresh токен
- `ParseToken(tokenStr, secret)` - парсит и проверяет JWT токен
- `IsTokenType(claims, tokenType)` - проверяет тип токена

## Ошибки

- `ErrInvalidSigningMethod` - возвращается при использовании неправильного метода подписи
- Стандартные ошибки из пакета `github.com/golang-jwt/jwt/v5`

## Безопасность

- Используется алгоритм HS256 для подписи токенов
- Обязательная проверка метода подписи при парсинге токена
- Рекомендуется использовать разные секретные ключи для access и refresh токенов
- Всегда устанавливайте разумное время жизни токенов

## Пример конфигурации

Рекомендуемые времена жизни токенов:
- Access токен: 15-30 минут
- Refresh токен: 7-30 дней