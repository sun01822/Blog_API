package response

import (
	"Blog_API/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// SuccessResponse creates a success response with a message, status code, and details.
func SuccessResponse(c echo.Context, message string, details interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": message,
		"details": details,
	})
}

func ErrorResponse(c echo.Context, err error, message string) error {
	statusCode := utils.StatusCode(err)
	return c.JSON(statusCode, map[string]interface{}{
		"error":   err.Error(),
		"message": message,
		"status":  statusCode,
	})
}
