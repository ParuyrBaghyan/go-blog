package utils

import "github.com/gin-gonic/gin"

func RegisterStaticFolder(server *gin.Engine) {
	server.Static("/media", "./media")
}
