package middleware

import (
	"issue-maintainer/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	header := ctx.Request.Header.Get("Authorization")

	if header == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization failed",
		})
		return
	}

	/* Auth header string is "Bearer token", expected length is 2 after splitting i.e. ["Bearer", "token"] */
	length := len(strings.Split(header, " "))

	if length != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid auth header format",
		})
		return
	}

	/* ["Bearer", "token"][1] is "token" */
	token := strings.Split(header, " ")[1]

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token is not valid",
		})
		return
	}

	userId, err := utils.VerifyJwtToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	/* can set any value in the context and retrieve it later inside the handler of certain routes in which this middleware is being used */
	ctx.Set("userId", userId)

	/* make sure we call ctx.Next() after checking things, to allow the request to go inside other middlewares or to the handler */
	ctx.Next()
}
