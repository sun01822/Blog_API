package routes

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type blogRoutes struct {
	echo           *echo.Echo
	blogController domain.BlogController
}

func NewBlogRoutes(e *echo.Echo, controller domain.BlogController) *blogRoutes {
	return &blogRoutes{
		echo:           e,
		blogController: controller,
	}
}

func (b *blogRoutes) InitBlogRoutes() {
	e := b.echo
	b.initBlogRoutes(e)
}

func (b *blogRoutes) initBlogRoutes(e *echo.Echo) {

	// group the routes
	common := e.Group("blog_api")
	version := common.Group("/v1")

	blog := version.Group("/blog")

	// blog routes
	blog.POST("/create", b.blogController.CreateBlogPost, middlewares.Auth)
	blog.GET("/get", b.blogController.GetBlogPost)
	blog.GET("/getAll", b.blogController.GetBlogPosts)
	blog.GET("/get/user", b.blogController.GetBlogPostsOfUser)
	blog.PUT("/update", b.blogController.UpdateBlogPost, middlewares.Auth)
	blog.DELETE("/delete", b.blogController.DeleteBlogPost, middlewares.Auth)

	// like and comment routes
	blog.POST("/reaction", b.blogController.AddAndRemoveReaction, middlewares.Auth)
	blog.POST("/comment", b.blogController.AddComment, middlewares.Auth)
	blog.GET("/comment", b.blogController.GetComments, middlewares.Auth)
	blog.DELETE("/comment", b.blogController.DeleteComment, middlewares.Auth)
	//blog.PUT("/comment/:userID/:postID/:commentID", b.blogController.UpdateComment, middlewares.Auth)

}
