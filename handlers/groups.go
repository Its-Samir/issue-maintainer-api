package handlers

import (
	"issue-maintainer/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateGroup(ctx *gin.Context) {
	var group models.Group
	err := ctx.ShouldBindJSON(&group)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse data",
		})
		return
	}

	/* get the userId we have set in the authentication middleware */
	userId := ctx.GetInt64("userId")

	/* assign the userId to the group's userId field, ensuring that it is the authenticated user */
	group.UserId = userId

	err = group.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create group",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Group created",
	})
}

func JoinGroup(ctx *gin.Context) {
	groupId := ctx.Param("id")
	parsedId, err := strconv.ParseInt(groupId, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the id",
		})
		return
	}

	userId := ctx.GetInt64("userId")

	err = models.AddUserIntoGroup(userId, parsedId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create group",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Group created",
	})
}

func DeleteGroup(ctx *gin.Context) {
	groupId := ctx.Param("id")
	parsedId, err := strconv.ParseInt(groupId, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse id",
		})
		return
	}

	group, notFoundErr := models.GetGroupById(parsedId)

	if notFoundErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Could not find group",
		})
		return
	}

	userId := ctx.GetInt64("userId")

	if group.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	err = models.DeleteGroupById(group.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete group",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Group deleted",
	})
}
