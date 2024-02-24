package routes

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type userRoutes struct {
	echo *echo.Echo
	userController domain.UserController
}

func NewUserRoutes(e *echo.Echo, controller domain.UserController) *userRoutes{
	return &userRoutes{
		echo: e,
		userController: controller,
	}
}

func (u *userRoutes) InitUserRoutes(){
	e:= u.echo
	u.initUserRoutes(e)
}

func (u *userRoutes) initUserRoutes(e *echo.Echo){

	// group the routes 
	user := e.Group("/blog_api/v1")

	// Login route
	user.POST("/user/login", u.userController.Login)

	user.POST("/user/create", u.userController.CreateUser)
	user.GET("/user/get/:id", u.userController.GetUser)
	user.GET("/user/get", u.userController.GetUsers)
	user.PUT("/user/update/:id", u.userController.UpdateUser, middlewares.Auth)
	user.DELETE("/user/delete/:id", u.userController.DeleteUser, middlewares.Auth)

}