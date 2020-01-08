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
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)
	u.CosmosUser = types.CosmosUser{
		Mnemonic: "",
	}

	// Save user
	h.DB.Create(&u)
	h.DB.Save(&u)

	u.Password = "" // Don't send password
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
	h.DB.Where(&types.User{
		Email: u.Email,
	}).First(&dbUser)

	if err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password)); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "Login failed")
	}

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
