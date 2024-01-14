package routes

import (
	"issue-maintainer/handlers"
	"issue-maintainer/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterGroupRoutes(server *gin.Engine) {
	router := server

	router.POST("/groups", middleware.Authenticate, handlers.CreateGroup)
	router.POST("/groups/:id", middleware.Authenticate, handlers.JoinGroup)
	router.POST("/groups/:id/issues", middleware.Authenticate, handlers.CreateIssue)
	router.PUT("/groups/:groupId/issues/:issueId", middleware.Authenticate, handlers.UpdateIssue)
	router.DELETE("/groups/:id", middleware.Authenticate, handlers.DeleteGroup)
}
