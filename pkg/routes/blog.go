package routes

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type blogRoutes struct {
	echo *echo.Echo
	foodController domain.BlogController
}

func NewBlogRoutes(e *echo.Echo, controller domain.BlogController) *blogRoutes{
	return &blogRoutes{
		echo: e,
		foodController: controller,
	}
}

func (b *blogRoutes) InitBlogRoutes(){
	e:= b.echo
	b.initBlogRoutes(e)
}

func (b *blogRoutes) initBlogRoutes(e *echo.Echo){
	
	// group the routes 
	blog := e.Group("/blog_api/v1")

	blog.POST("/blog/create", b.foodController.CreateBlogPost, middlewares.Auth)
	blog.GET("/blog/get/:id", b.foodController.GetBlogPost)
	blog.GET("/blog/get/user/:userID", b.foodController.GetBlogPosts)
	blog.PUT("/blog/update/:id", b.foodController.UpdateBlogPost, middlewares.Auth)
	blog.DELETE("/blog/delete/:id", b.foodController.DeleteBlogPost, middlewares.Auth)

}
