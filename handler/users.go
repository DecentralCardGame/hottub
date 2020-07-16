package handler

import (
	"github.com/labstack/echo/v4"
	"hottub/types"
	"hottub/utils"
	"net/http"
	"strconv"
)

// Get all users
// @Summary Get all users
// @Description Get all users
// @ID getUsers
// @Accept  json
// @Produce  json
// @Success 200 {object} types.PublicUsersRepsonse "ok"
// @Router /users [get]
func (h *Handler) GetUsers(c echo.Context) error {
	var users []types.User
	h.DB.Find(&users)
	return c.JSON(http.StatusOK, types.NewPublicUsersResponse(users))
}

// Get User
// @Summary Get a given User by ID
// @Description Get a given User by ID
// @ID getUser
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Success 200 {object} types.PublicUserResponse	"ok"
// @Router /users/{id} [get]
func (h *Handler) GetUsersById(c echo.Context) error {
	var user types.User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(utils.ErrorParameterNotInteger.Status, utils.ErrorParameterNotInteger)
	}
	h.DB.First(&user, id)
	return c.JSON(http.StatusOK, types.NewPublicUserResponse(&user))
}

// Create User
// @Summary Create a user
// @Description Create a user
// @ID createUser
// @Accept  json
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Param email body string true "E-Mail"
// @Param mnemonic body string true "Mnemonic"
// @Success 200 {object} types.PublicUserResponse	"ok"
// @Router /users [post]
func (h *Handler) CreateUser(c echo.Context) error {
	user := new(types.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(utils.ErrorCannotParseFields.Status, utils.ErrorCannotParseFields)
	}

	h.DB.NewRecord(user)
	h.DB.Create(&user)

	return c.JSON(http.StatusOK, types.NewPublicUserResponse(user))
}

// Update User
// @Summary Updates a given User by ID
// @Description Updates a given User by ID
// @ID updateUser
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Success 200 {object} types.PublicUserResponse	"ok"
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(c echo.Context) error {
	reqUser := new(types.User)
	var user types.User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(utils.ErrorParameterNotInteger.Status, utils.ErrorParameterNotInteger)
	}

	if err = c.Bind(reqUser); err != nil {
		return c.JSON(utils.ErrorCannotParseFields.Status, utils.ErrorCannotParseFields)
	}

	h.DB.First(&user, id)
	h.DB.NewRecord(reqUser)
	h.DB.Create(&reqUser)

	return c.JSON(http.StatusOK, types.NewPublicUserResponse(&user))
}

// Delete User
// @Summary Deletes a given User by ID
// @Description Deletes a given User by ID
// @ID deleteUser
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Success 200 {string} string	"ok"
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	h.DB.Where("id = ?", id).Delete(&types.User{})
	return c.String(http.StatusOK, "ok")
}
