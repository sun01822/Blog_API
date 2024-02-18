package main

import (
	"Blog_API/pkg/containers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e:= echo.New()
	e.Use(middleware.CORS())
	containers.Serve(e)
}