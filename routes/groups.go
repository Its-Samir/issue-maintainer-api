package routes

import (
	"issue-maintainer/handlers"
	"issue-maintainer/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterGroupRoutes(server *gin.Engine) {
	router := server

	router.POST("/groups", middleware.Authenticate, handlers.CreateGroup)
	router.DELETE("/groups/:id", middleware.Authenticate, handlers.DeleteGroup)
}
