package middlewares

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	godotenv.Load()
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		SigningMethod: "HS256",
	})
}

func CreateToken(id int, role string) (string, error) {
	godotenv.Load()
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ExtractToken(e echo.Context) (int, string) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		Id := int(claims["id"].(float64))
		Role := claims["role"].(string)
		return Id, Role
	}
	return 0, ""
}

func ExtractRole(c echo.Context) (string, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role, ok := claims["role"].(string)
	if !ok {
		return "", errors.New("role not found in token claims")
	}
	return role, nil
}
