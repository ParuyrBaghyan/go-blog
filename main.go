package main

import (
	"github.com/gin-gonic/gin"
	"go-blog/db"
	"go-blog/routes"
	"go-blog/utils"
)

func main() {
	db.InitDB()

	server := gin.Default()

	utils.RegisterStaticFolder(server)

	routes.RegisterRouters(server)

	defer server.Run(":8080")
}
