package hottub

import (
	"bytes"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	db "hottub/db"
	types "hottub/types"
	"log"
	"net/http"
	"time"
)

// THIS IS FOR TESTING PURPOSES ONLY
// DO NOT USE IN PRODUCTION!!!
var secretkey string = "test"

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user types.User
	db.Manager().Where(&types.User{Username: username}).First(&user)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	if !(bytes.Equal(hash, []byte(user.Password))) {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = &user
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
