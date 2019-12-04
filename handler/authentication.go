package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"hottub/types"
	"net/http"
	"time"
)

func (h *Handler) Register(c echo.Context) (err error) {
	// Bind
	u := &types.User{}
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	u.Password = string(hash)

	// Save user
	h.DB.NewRecord(u)
	h.DB.Create(&u)

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(types.User)
	if err = c.Bind(u); err != nil {
		return
	}

	// Find user
	var dbUser types.User
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	h.DB.Where(&types.User{
		Username: u.Username,
		Password: string(hash),
	}).First(&dbUser)

	// Equation logic

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	u.Password = "" // Don't send password
	return c.JSON(http.StatusOK, u)
}
