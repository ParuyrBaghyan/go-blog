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
	authProfile.GET("/", getUserProfile)
	authProfile.PUT("/update", updateUser)
	authProfile.DELETE("/delete", deleteUser)

	server.GET("/posts", getAllPosts)
	server.GET("/posts/:id", getPost)
	authPosts := server.Group("/posts")
	authPosts.Use(middlewares.Authenticate)
	authPosts.POST("/create", createPost)
	authPosts.PUT("/update/:id", updatePost)
	authPosts.DELETE("/delete/:id", deletePost)
}
