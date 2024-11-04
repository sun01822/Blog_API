package containers

import (
	"Blog_API/pkg/config"
	"Blog_API/pkg/connection"
	"Blog_API/pkg/controllers"
	"Blog_API/pkg/repositories"
	"Blog_API/pkg/routes"
	"Blog_API/pkg/services"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

// Serve is a function that returns a new instance of echo.Echo
func Serve(e *echo.Echo) {

	// Config initialization
	config.SetConfig()

	// Database initialization
	db := connection.GetDB()

	// Repository initialization
	userRepo := repositories.NewUserRepo(db)
	blogRepo := repositories.NewBlogRepo(db)

	// Service initialization
	userService := services.SetUserService(userRepo)
	blogService := services.NewBlogService(blogRepo, userRepo)

	// Controller initialization
	userController := controllers.SetUserController(userService)
	blogController := controllers.NewBlogController(blogService, userService)

	user := routes.NewUserRoutes(e, userController)
	user.InitUserRoutes()
	blog := routes.NewBlogRoutes(e, blogController)
	blog.InitBlogRoutes()

	// Starting Server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}
