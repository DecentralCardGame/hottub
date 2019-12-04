package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) MainRoute(c echo.Context) error {
	return c.String(http.StatusOK, "DecentralCardGame - Hottub")
}
