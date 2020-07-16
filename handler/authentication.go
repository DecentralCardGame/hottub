package handler

import (
	"github.com/labstack/echo"
	"hottub/types"
	"hottub/utils"
	"net/http"
)

func (h *Handler) Register(c echo.Context) (err error) {
	var u types.User
	req := &types.UserRegisterRequest{}

	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	// Save user
	// TODO: Move logic to store
	h.DB.Create(&u)
	h.DB.Save(&u)

	return c.JSON(http.StatusOK, types.NewUserLoginResponse(&u))
}

func (h *Handler) Login(c echo.Context) (err error) {
	req := &types.UserLoginRequest{}

	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	// Find user
	// TODO: Move logic to store
	var dbUser types.User
	h.DB.Where(&types.User{
		Username: req.Username,
	}).First(&dbUser)

	if !dbUser.ComparePassword(req.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	return c.JSON(http.StatusOK, types.NewUserLoginResponse(&dbUser))
}
