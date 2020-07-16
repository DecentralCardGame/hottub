package handler

import (
	"github.com/labstack/echo/v4"
	"hottub/types"
	"net/http"
)

// Register
// @Summary Information
// @Description Gives back Information about the API
// @ID root
// @Accept  json
// @Produce  json
// @Success 200 {object} types.WelcomeResponse "ok"
// @Router / [get]
func (h *Handler) MainRoute(c echo.Context) error {
	return c.JSON(http.StatusOK, types.NewWelcomeResponse())
}
