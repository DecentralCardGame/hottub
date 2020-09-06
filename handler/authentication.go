package handler

import (
	"github.com/labstack/echo/v4"
	"hottub/types"
	"hottub/utils"
	"net/http"
)

// Register
// @Summary Register a new user
// @Description Register a new user
// @ID register
// @Accept  json
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Param email body string true "E-Mail"
// @Param mnemonic body string true "Mnemonic"
// @Success 200 {object} types.UserLoginResponse	"ok"
// @Router /register [post]
func (h *Handler) Register(c echo.Context) (err error) {
	var u types.User
	req := &types.UserRegisterRequest{}

	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := h.UserStore.CreateNewUser(&u); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, types.NewUserLoginResponse(&u))
}

// Login
// @Summary Login as an existing user
// @Description Login as an existing user
// @ID login
// @Accept  json
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} types.UserLoginResponse	"ok"
// @Router /login [post]
func (h *Handler) Login(c echo.Context) (err error) {
	req := &types.UserLoginRequest{}

	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	dbUser, err := h.UserStore.GetByUsername(req.Username)

	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewError(err))
	}

	if !dbUser.ComparePassword(req.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}

	return c.JSON(http.StatusOK, types.NewUserLoginResponse(dbUser))
}
