package handler

import (
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"hottub/types"
	"hottub/utils"
	"net/http"
)

func (h *Handler) Register(c echo.Context) (err error) {
	// Bind
	u := &types.User{}

	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Username == "" || u.Email == "" || u.Password == "" || u.Mnemonic == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid credentials")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)

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
		Username: u.Username,
	}).First(&dbUser)

	if err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password)); err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "Login failed")
	}

	dbUser.Token = utils.GenerateJWT(u.ID)

	dbUser.Password = "" // Don't send password
	return c.JSON(http.StatusOK, dbUser)
}
