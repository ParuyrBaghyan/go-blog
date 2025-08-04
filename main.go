package main

import (
	"go-blog/db"
	"go-blog/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRouters(server)

	server.Run(":8080")
}
