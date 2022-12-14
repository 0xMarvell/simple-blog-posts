package routes

import (
	"github.com/0xMarvell/simple-blog-posts/pkg/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterBlogRoutes handles all routing for the API.
func RegisterBlogRoutes(r *gin.Engine) {
	r.POST("/api/v1/posts", controllers.CreatePost)
	r.GET("/api/v1/posts", controllers.GetPosts)
	r.GET("/api/v1/posts/:id", controllers.GetPost)
	r.PUT("/api/v1/posts/:id", controllers.UpdatePost)
	r.DELETE("/api/v1/posts/:id", controllers.DeletePost)
}
