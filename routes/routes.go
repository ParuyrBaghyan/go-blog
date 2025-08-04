package routes

import (
	"go-blog/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/signin", signin)

	authenticate := server.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.GET("/profile", getUserProfile)

}
