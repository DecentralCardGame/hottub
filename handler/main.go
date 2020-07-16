package handler

import (
	"github.com/labstack/echo/v4"
	"hottub/types"
	"net/http"
)

func (h *Handler) MainRoute(c echo.Context) error {
	return c.JSON(http.StatusOK, types.NewWelcomeResponse())
}
