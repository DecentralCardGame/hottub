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
	isAdmin, err := h.UserStore.CheckUserAdmin(utils.GetUserIDFromContext(c))

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	if !isAdmin {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}

	return c.JSON(http.StatusOK, types.NewPublicUsersResponse(h.UserStore.GetAllUsers()))
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

	isAdminOrMe, err := h.UserStore.CheckUserAdminOrMe(utils.GetUserIDFromContext(c), id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	if !isAdminOrMe {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
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
	var u types.User
	req := &types.UserRegisterRequest{}

	isAdmin, err := h.UserStore.CheckUserAdmin(utils.GetUserIDFromContext(c))

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	if !isAdmin {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}

	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := h.UserStore.CreateNewUser(&u); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, types.NewPublicUserResponse(&u))
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
	reqUser := &types.User{}
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(utils.ErrorParameterNotInteger.Status, utils.ErrorParameterNotInteger)
	}

	isAdminOrMe, err := h.UserStore.CheckUserAdminOrMe(utils.GetUserIDFromContext(c), id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	if !isAdminOrMe {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}

	if err = c.Bind(reqUser); err != nil {
		return c.JSON(utils.ErrorCannotParseFields.Status, utils.ErrorCannotParseFields)
	}

	h.UserStore.UpdateUser(reqUser)

	return c.JSON(http.StatusOK, types.NewPublicUserResponse(reqUser))
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
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(utils.ErrorParameterNotInteger.Status, utils.ErrorParameterNotInteger)
	}

	isAdmin, err := h.UserStore.CheckUserAdmin(utils.GetUserIDFromContext(c))
	isMe, err := h.UserStore.CheckUserMe(utils.GetUserIDFromContext(c), id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(err))
	}

	if !isAdmin && !isMe {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}

	err = h.UserStore.DeleteUser(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.String(http.StatusOK, "ok")
}
