package main

import (
	_ "Blog_API/docs"
	"Blog_API/pkg/containers"
	"Blog_API/pkg/utils/response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	m "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
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
// @host 192.168.10.135:8080
// @BasePath /blog_api/v1
func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(m.RemoveTrailingSlash())
	e.Use(m.LoggerWithConfig(m.LoggerConfig{
		Format:           `${time_custom} ${remote_ip} ${host} ${method} ${uri} ${status} ${latency_human} ${bytes_in} ${bytes_out} "${user_agent}"` + "\n",
		CustomTimeFormat: "2006-01-02T15:04:05.00",
	}))
	e.Use(m.Secure())
	e.Use(m.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/h3a1t5", func(c echo.Context) error {
		return response.SuccessResponse(c, "Health Check", "Server is running")
	})
	containers.Serve(e)
}
