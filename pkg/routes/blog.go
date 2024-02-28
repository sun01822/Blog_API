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
	blog.GET("/blog/get/:postID", b.blogController.GetBlogPost)
	blog.GET("/blog/get", b.blogController.GetBlogPosts)
	blog.GET("/blog/get/user/:userID", b.blogController.GetBlogPostsOfUser)
	blog.PUT("/blog/update/:userID/:postID", b.blogController.UpdateBlogPost, middlewares.Auth)
	blog.DELETE("/blog/delete/:userID/:postID", b.blogController.DeleteBlogPost, middlewares.Auth)

	// like and comment routes
	blog.POST("/blog/like/:userID/:postID", b.blogController.AddAndRemoveLike, middlewares.Auth)
	blog.POST("/blog/comment/:userID/:postID", b.blogController.AddComment, middlewares.Auth)
	blog.GET("/blog/comment/:postID", b.blogController.GetComments)
	blog.GET("/blog/comment/:postID/:commentID", b.blogController.GetCommentByUserID)
	blog.DELETE("/blog/comment/:userID/:postID/:commentID", b.blogController.DeleteComment, middlewares.Auth)
	blog.PUT("/blog/comment/:userID/:postID/:commentID", b.blogController.UpdateComment, middlewares.Auth)
	
}
