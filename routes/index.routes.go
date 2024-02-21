package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexHandle(c echo.Context) error {
	return c.String(http.StatusOK, "REST API Gorm")
}
