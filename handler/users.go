package handler

import (
	"github.com/labstack/echo"
	"hottub/db"
	"hottub/types"
	"net/http"
	"strconv"
)

func (h *Handler) GetUsers(c echo.Context) error {
	var users []types.User
	db.Manager().Find(&users)
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUsersById(c echo.Context) error {
	var user types.User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(types.ErrorParameterNotInteger.Status, types.ErrorParameterNotInteger)
	}
	db.Manager().First(&user, id)
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) CreateUser(c echo.Context) error {
	user := new(types.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(types.ErrorCannotParseFields.Status, types.ErrorCannotParseFields)
	}

	db.Manager().NewRecord(user)
	db.Manager().Create(&user)

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	reqUser := new(types.User)
	var user types.User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(types.ErrorParameterNotInteger.Status, types.ErrorParameterNotInteger)
	}

	if err = c.Bind(reqUser); err != nil {
		return c.JSON(types.ErrorCannotParseFields.Status, types.ErrorCannotParseFields)
	}

	db.Manager().First(&user, id)
	db.Manager().NewRecord(reqUser)
	db.Manager().Create(&reqUser)

	return c.JSON(http.StatusOK, reqUser)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	db.Manager().Where("id = ?", id).Delete(&types.User{})
	return c.String(http.StatusOK, "ok")
}
