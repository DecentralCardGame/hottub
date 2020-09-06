package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"time"
)

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString([]byte("secret"))
	return t
}

func GetClaimsFromContext(c echo.Context) jwt.MapClaims {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(jwt.MapClaims)
}

func GetUserIDFromContext(c echo.Context) int {
	claims := GetClaimsFromContext(c)
	id := int(claims["id"].(float64))
	return id
}
