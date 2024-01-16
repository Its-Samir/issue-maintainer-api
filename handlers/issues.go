package handlers

import (
	"issue-maintainer/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFeaturedIssues(ctx *gin.Context) {
	issues, err := models.GetFeaturedIssues()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get the issues",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"issues": issues,
	})
}

func GetIssues(ctx *gin.Context) {
	groupId := ctx.Param("groupId")
	parsedGroupId, groupIdParsedErr := strconv.ParseInt(groupId, 10, 64)

	if groupIdParsedErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse issueId",
		})
		return
	}

	issues, err := models.GetAllIssues(parsedGroupId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get the issues",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"issues": issues,
	})
}

func GetIssue(ctx *gin.Context) {
	issueId := ctx.Param("issueId")
	parsedIssueId, issueIdParsedErr := strconv.ParseInt(issueId, 10, 64)

	if issueIdParsedErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse issueId",
		})
		return
	}

	groupId := ctx.Param("groupId")
	parsedGroupId, groupIdParsedErr := strconv.ParseInt(groupId, 10, 64)

	if groupIdParsedErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse issueId",
		})
		return
	}

	issue, notFoundErr := models.GetIssueById(parsedIssueId, parsedGroupId)

	if notFoundErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find issue",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"issues": issue,
	})
}

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

func UpdateIssue(ctx *gin.Context) {
	var updatedIssue models.Issue

	err := ctx.ShouldBindJSON(&updatedIssue)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the data",
		})
		return
	}

	issueId := ctx.Param("issueId")
	parsedIssueId, issueIdParsedErr := strconv.ParseInt(issueId, 10, 64)

	if issueIdParsedErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse issueId",
		})
		return
	}

	groupId := ctx.Param("groupId")

	parsedGroupId, groupIdParsedErr := strconv.ParseInt(groupId, 10, 64)

	if groupIdParsedErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse groupId",
		})
		return
	}

	issue, notFoundErr := models.GetIssueById(parsedIssueId, parsedGroupId)

	if notFoundErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find issue",
		})
		return
	}

	userId := ctx.GetInt64("userId")

	if parsedGroupId != issue.GroupId || userId != issue.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}

	err = models.UpdateIssueById(issue.Id, updatedIssue)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update issue",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Issue updated",
	})
}
