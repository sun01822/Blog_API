package routes

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type userRoutes struct {
	echo *echo.Echo
	foodController domain.UserController
}

func NewUserRoutes(e *echo.Echo, controller domain.UserController) *userRoutes{
	return &userRoutes{
		echo: e,
		foodController: controller,
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
	user.POST("/user/login", u.foodController.Login)

	user.POST("/user/create", u.foodController.CreateUser)
	user.GET("/user/get/:id", u.foodController.GetUser)
	user.GET("/user/get", u.foodController.GetUsers)
	user.PUT("/user/update/:id", u.foodController.UpdateUser, middlewares.Auth)
	user.DELETE("/user/delete/:id", u.foodController.DeleteUser, middlewares.Auth)

}