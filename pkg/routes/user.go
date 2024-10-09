package routes

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type userRoutes struct {
	echo           *echo.Echo
	userController domain.UserController
}

func NewUserRoutes(e *echo.Echo, controller domain.UserController) *userRoutes {
	return &userRoutes{
		echo:           e,
		userController: controller,
	}
}

func (u *userRoutes) InitUserRoutes() {
	e := u.echo
	u.initUserRoutes(e)
}

func (u *userRoutes) initUserRoutes(e *echo.Echo) {

	// group the routes
	common := e.Group("blog_api")
	version := common.Group("/v1")

	user := version.Group("/user")

	// Login route
	user.POST("/login", u.userController.Login)

	user.POST("/create", u.userController.CreateUser)
	user.GET("/get", u.userController.GetUser)
	user.GET("/get/users", u.userController.GetUsers)
	user.PUT("/update/:userID", u.userController.UpdateUser, middlewares.Auth)
	user.DELETE("/delete/:userID", u.userController.DeleteUser, middlewares.Auth)

}
