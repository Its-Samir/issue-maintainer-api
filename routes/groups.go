package routes

import (
	"issue-maintainer/handlers"
	"issue-maintainer/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterGroupRoutes(server *gin.Engine) {
	router := server

	router.GET("/issues", handlers.GetFeaturedIssues)

	router.POST("/groups", middleware.Authenticate, handlers.CreateGroup)

	router.POST("/groups/:id", middleware.Authenticate, handlers.JoinGroup)

	router.GET("/groups/:groupId/issues", handlers.GetIssues)

	router.GET("/groups/:groupId/issues/:issueId", handlers.GetIssue)

	router.POST("/groups/:id/issues", middleware.Authenticate, handlers.CreateIssue)

	router.PUT("/groups/:groupId/issues/:issueId", middleware.Authenticate, handlers.UpdateIssue)

	router.DELETE("/groups/:groupId/issues/:issueId", middleware.Authenticate, handlers.DeleteIssue)

	router.DELETE("/groups/:id", middleware.Authenticate, handlers.DeleteGroup)
}
