package handlers

import (
	"issue-maintainer/models"
	"issue-maintainer/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse data",
		})
		return
	}

	err = user.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not register user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registration successfull",
	})
}

func Login(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse data",
		})
		return
	}

	id, validationErr := user.ValidateCredentials()

	if validationErr != nil {
		ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	token, tokenErr := utils.GenerateJwtToken(user.Email, id)

	if tokenErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authentication failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login successfull",
		"token":   token,
	})
}
