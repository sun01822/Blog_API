package routes

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type UserRoutes struct {
	echo           *echo.Echo
	userController domain.Controller
}

func NewUserRoutes(e *echo.Echo, controller domain.Controller) UserRoutes {
	return UserRoutes{
		echo:           e,
		userController: controller,
	}
}

func (u *UserRoutes) InitUserRoutes() {
	e := u.echo
	u.initUserRoutes(e)
}

func (u *UserRoutes) initUserRoutes(e *echo.Echo) {

	// group the routes
	common := e.Group("blog_api")
	version := common.Group("/v1")

	user := version.Group("/user")

	// Login route
	user.POST("/login", u.userController.Login)
	user.POST("/logout", u.userController.Logout, middlewares.Auth)

	user.POST("/create", u.userController.CreateUser)
	user.GET("/get", u.userController.GetUser)
	user.GET("/getAll", u.userController.GetUsers)
	user.PUT("/update", u.userController.UpdateUser, middlewares.Auth)
	user.DELETE("/delete", u.userController.DeleteUser, middlewares.Auth)
}
