package errorhandlers

import (
	"net/http"

	"github.com/labstack/echo"
	"go.mod/models"
)

func BadRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, models.Response{
		Status:  "400",
		Message: message,
	})
}

func OKRequest(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, models.Response{
		Status:  "200",
		Message: message,
	})
}
