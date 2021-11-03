package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
