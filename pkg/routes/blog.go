package routes

import (
	"Blog_API/pkg/domain"
	"Blog_API/pkg/middlewares"
	"github.com/labstack/echo/v4"
)

type blogRoutes struct {
	echo *echo.Echo
	blogController domain.BlogController
}

func NewBlogRoutes(e *echo.Echo, controller domain.BlogController) *blogRoutes{
	return &blogRoutes{
		echo: e,
		blogController: controller,
	}
}

func (b *blogRoutes) InitBlogRoutes(){
	e:= b.echo
	b.initBlogRoutes(e)
}

func (b *blogRoutes) initBlogRoutes(e *echo.Echo){
	
	// group the routes 
	blog := e.Group("/blog_api/v1")

	// blog routes
	blog.POST("/blog/create/:userID", b.blogController.CreateBlogPost, middlewares.Auth)
	blog.GET("/blog/get/:id", b.blogController.GetBlogPost)
	blog.GET("/blog/get", b.blogController.GetBlogPosts)
	blog.GET("/blog/get/user/:userID", b.blogController.GetBlogPostsOfUser)
	blog.PUT("/blog/update/:userID/:id", b.blogController.UpdateBlogPost, middlewares.Auth)
	blog.DELETE("/blog/delete/:userID/:id", b.blogController.DeleteBlogPost, middlewares.Auth)

	// like and comment routes
	blog.POST("/blog/like/:userID/:id", b.blogController.AddAndRemoveLike, middlewares.Auth)
	blog.POST("/blog/comment/:userID/:id", b.blogController.AddComment, middlewares.Auth)
	blog.GET("/blog/comment/:id", b.blogController.GetComments)
	blog.GET("/blog/comment/:id/:commentID", b.blogController.GetCommentByUserID)
	blog.DELETE("/blog/comment/:userID/:id/:commentID", b.blogController.DeleteComment, middlewares.Auth)
	blog.PUT("/blog/comment/:userID/:id/:commentID", b.blogController.UpdateComment, middlewares.Auth)
	
}
