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
func Serve(e *echo.Echo)  {

	// Config initalization
	config.SetConfig()


	// Database initalization
	db:= connection.GetDB()

	// Repository initalization
	userRepo := repositories.NewUserRepo(db)
	blogRepo := repositories.NewBlogRepo(db)

	// Service initalization
	userService := services.NewUserService(userRepo)
	blogService := services.NewBlogService(blogRepo)

	// Controller initalization
	userController := controllers.NewUserController(userService)
	blogController := controllers.NewBlogController(blogService,userService)

	// Routes
	user := routes.NewUserRoutes(e, userController)
	user.InitUserRoutes()
	blog := routes.NewBlogRoutes(e, blogController)
	blog.InitBlogRoutes()



	// Starting Server
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.LocalConfig.Port)))
}