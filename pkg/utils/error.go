package utils

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

// StatusCode returns the appropriate HTTP status code based on the error type.
func StatusCode(err error) int {
	var he *echo.HTTPError
	if errors.As(err, &he) {
		return he.Code
	}
	return http.StatusInternalServerError
}
