package handlers

import (
	"issue-maintainer/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIssues(ctx *gin.Context) {}

func GetIssue(ctx *gin.Context) {}

func CreateIssue(ctx *gin.Context) {
	var issue models.Issue
	err := ctx.ShouldBindJSON(&issue)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the data",
		})
		return
	}

	userId := ctx.GetInt64("userId")
	groupId := ctx.Param("id")
	parsedGroupId, err := strconv.ParseInt(groupId, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	issue.UserId = userId
	issue.GroupId = parsedGroupId

	err = issue.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create issue",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Issue created",
	})
}

func UpdateIssue(ctx *gin.Context) {}
