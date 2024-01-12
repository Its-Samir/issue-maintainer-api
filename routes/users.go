package routes

import (
	"issue-maintainer/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	/* just for naming convention */
	router := server

	router.POST("/signup", handlers.Signup)

	router.POST("/login", handlers.Login)
}
