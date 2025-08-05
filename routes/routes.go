package routes

import (
	"go-blog/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/signin", signin)

	authProfile := server.Group("/profile")
	authProfile.Use(middlewares.Authenticate)
	authProfile.GET("/profile", getUserProfile)
	authProfile.PUT("/profile/update", updateUser)
	authProfile.DELETE("/profile/delete", deleteUser)

	server.GET("/posts", getAllPosts)
	server.GET("/posts/:id", getPost)
	server.POST("/posts/createPost", createPost)
	// authPosts := server.Group("/posts")
	// authPosts.Use(middlewares.Authenticate)
	// authPosts.POST("/posts/createPost", createPost)
}
