package main

import (
	"issue-maintainer/db"
	"issue-maintainer/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterUserRoutes(server)

	server.Run(":8080")
}
