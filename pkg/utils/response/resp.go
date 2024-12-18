package response

import (
	"Blog_API/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SuccessRes struct {
	Message string      `json:"message"`
	Details interface{} `json:"details"`
	Status  int         `json:"status"`
}

type ErrorRes struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Status  int    `json:"status"`
}

// SuccessResponse creates a success response with a message, status code, and details.
func SuccessResponse(c echo.Context, message string, details interface{}) error {
	return c.JSON(http.StatusOK, SuccessRes{
		Message: message,
		Details: details,
		Status:  http.StatusOK,
	})
}

func ErrorResponse(c echo.Context, err error, message string) error {
	statusCode := utils.StatusCode(err)
	return c.JSON(statusCode, ErrorRes{
		Message: message,
		Error:   err.Error(),
		Status:  statusCode,
	})
}

// ErrorResponseWithStatus creates an error response with a message, status code, and details.
func ErrorResponseWithStatus(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, ErrorRes{
		Error:   http.StatusText(statusCode),
		Message: message,
		Status:  statusCode,
	})
}
