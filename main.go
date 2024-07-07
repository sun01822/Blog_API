package main

import (
	"Blog_API/pkg/containers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "Blog_API/docs"
)


// swagger docs
// @title Blog API
// @version 1.0
// @description This is a sample server for Blog CRUD Operation
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /blog_api/v1
func main() {
	e:= echo.New()
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	containers.Serve(e)
}


