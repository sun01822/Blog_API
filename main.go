package main

import (
	"Blog_API/pkg/containers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "Blog_API/docs"
)



func main() {
	e:= echo.New()
	e.Use(middleware.CORS())

	e.GET("/health", HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	containers.Serve(e)
}

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
	   "data": "Server is running",
	})
 }

