package routes

import "github.com/gin-gonic/gin"

func RegisterRouters(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/signin", signin)
}
